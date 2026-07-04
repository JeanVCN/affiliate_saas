package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	httpapi "github.com/JeanVCN/affiliate_saas/backend/internal/http"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestFirstSliceHTTPFlowWithPostgres(t *testing.T) {
	databaseURL := os.Getenv("AFFILIATE_TEST_DATABASE_URL")
	if databaseURL == "" {
		t.Skip("set AFFILIATE_TEST_DATABASE_URL to run PostgreSQL integration tests")
	}

	ctx := context.Background()
	pool := openIsolatedPool(t, ctx, databaseURL)
	defer pool.Close()

	runMigrations(t, ctx, pool)

	router := httpapi.NewRouter(httpapi.Dependencies{
		AppEnv: "test",
		DB:     pool,
	})

	client := newTestClient(router)
	signup := postJSON[authResponse](t, client, "/api/v1/auth/signup", `{
		"email":"founder@example.com",
		"password":"CorrectHorse123",
		"display_name":"Founder",
		"workspace_name":"Creator Lab"
	}`)
	if len(signup.Workspaces) != 1 || signup.Workspaces[0].WorkspaceID == "" {
		t.Fatalf("signup workspaces = %+v", signup.Workspaces)
	}

	workspaceID := signup.Workspaces[0].WorkspaceID
	program := postJSON[workspaceProgramResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/programs", `{
		"marketplace_name":"TikTok Shop",
		"program_name":"TikTok Shop Affiliate",
		"external_account_label":"main"
	}`)
	product := postJSON[productResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/products", `{
		"name":"Creator Camera",
		"category":"electronics"
	}`)
	campaign := postJSON[campaignResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns", fmt.Sprintf(`{
		"product_id":%q,
		"name":"Creator Camera Launch"
	}`, product.ID))
	channelPackage := postJSON[channelPackageResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns/"+campaign.ID+"/channel-packages", `{
		"channel":"TikTok",
		"title":"Camera review",
		"body":"Show the creator setup and disclose the affiliate link."
	}`)
	campaignDetail := getJSON[campaignResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns/"+campaign.ID)
	if campaignDetail.ProductID != product.ID || len(campaignDetail.ChannelPackages) != 1 || campaignDetail.ChannelPackages[0].ID != channelPackage.ID {
		t.Fatalf("campaign detail = %+v", campaignDetail)
	}
	task := postJSON[publishingTaskResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns/"+campaign.ID+"/publishing-tasks", fmt.Sprintf(`{
		"channel_package_id":%q,
		"channel":"tiktok",
		"title":"Publish camera review",
		"notes":"Manual publish only"
	}`, channelPackage.ID))
	updatedTask := patchJSON[publishingTaskResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns/"+campaign.ID+"/publishing-tasks/"+task.ID, `{"status":"done"}`)
	if updatedTask.Status != "done" || updatedTask.CompletedAt == "" {
		t.Fatalf("updated publishing task = %+v", updatedTask)
	}
	tasks := getJSON[[]publishingTaskResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns/"+campaign.ID+"/publishing-tasks")
	if len(tasks) != 1 || tasks[0].ID != task.ID {
		t.Fatalf("publishing tasks = %+v", tasks)
	}
	updatedCampaign := patchJSON[campaignResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns/"+campaign.ID, `{"status":"ready"}`)
	if updatedCampaign.Status != "ready" {
		t.Fatalf("campaign status = %q, want ready", updatedCampaign.Status)
	}
	complianceCheck := postJSON[complianceCheckResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns/"+campaign.ID+"/compliance-checks", `{
		"channel":"tiktok",
		"title":"Camera review",
		"body":"Publi com link afiliado. Show the creator setup."
	}`)
	if len(complianceCheck.Findings) != 1 || complianceCheck.Findings[0].Code != "basic_check_passed" {
		t.Fatalf("compliance check = %+v", complianceCheck)
	}
	campaigns := getJSON[[]campaignResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/campaigns")
	if len(campaigns) != 1 || campaigns[0].ID != campaign.ID {
		t.Fatalf("campaigns = %+v", campaigns)
	}
	offer := postJSON[offerResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/products/"+product.ID+"/offers", fmt.Sprintf(`{
		"workspace_program_id":%q,
		"title":"TikTok Shop offer",
		"price_cents":129900,
		"currency":"brl"
	}`, program.ID))
	link := postJSON[linkResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/links", fmt.Sprintf(`{
		"product_id":%q,
		"offer_id":%q,
		"destination_url":"https://example.com/products/creator-camera",
		"label":"TikTok bio"
	}`, product.ID, offer.ID))
	conversionImport := postJSON[conversionImportResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/conversion-imports", `{"source":"manual"}`)
	conversionRow := postJSON[conversionImportRowResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/conversion-imports/"+conversionImport.ID+"/rows", fmt.Sprintf(`{
		"product_id":%q,
		"affiliate_link_id":%q,
		"order_reference":"order-1001",
		"gross_amount_cents":129900,
		"commission_cents":12990,
		"currency":"brl",
		"raw_payload":{"source":"manual-fixture"}
	}`, product.ID, link.ID))
	conversionDetail := getJSON[conversionImportResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/conversion-imports/"+conversionImport.ID)
	if conversionDetail.Source != "manual" || len(conversionDetail.Rows) != 1 || conversionDetail.Rows[0].ID != conversionRow.ID {
		t.Fatalf("conversion import detail = %+v", conversionDetail)
	}
	shortLink := postJSON[shortLinkResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/links/"+link.ID+"/short-links", `{"slug":"creator-camera"}`)

	redirectReq := httptest.NewRequest(http.MethodGet, "/r/"+shortLink.Slug, nil)
	redirectReq.Header.Set("Referer", "https://social.example/post")
	redirectReq.Header.Set("User-Agent", "integration-test")
	redirectRec := httptest.NewRecorder()
	router.ServeHTTP(redirectRec, redirectReq)

	if redirectRec.Code != http.StatusFound {
		t.Fatalf("redirect status = %d, want %d, body = %s", redirectRec.Code, http.StatusFound, redirectRec.Body.String())
	}
	if location := redirectRec.Header().Get("Location"); location != "https://example.com/products/creator-camera" {
		t.Fatalf("Location = %q", location)
	}

	clicksByProduct := getJSON[clickMetricsResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/analytics/clicks?group_by=product")
	if clicksByProduct.GroupBy != "product" || len(clicksByProduct.Items) != 1 {
		t.Fatalf("product metrics = %+v", clicksByProduct)
	}
	if clicksByProduct.Items[0].GroupID != product.ID || clicksByProduct.Items[0].Clicks != 1 {
		t.Fatalf("product metric item = %+v", clicksByProduct.Items[0])
	}

	clicksByLink := getJSON[clickMetricsResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/analytics/clicks?group_by=link")
	if clicksByLink.GroupBy != "link" || len(clicksByLink.Items) != 1 {
		t.Fatalf("link metrics = %+v", clicksByLink)
	}
	if clicksByLink.Items[0].GroupID != link.ID || clicksByLink.Items[0].Clicks != 1 {
		t.Fatalf("link metric item = %+v", clicksByLink.Items[0])
	}

	overview := getJSON[analyticsOverviewResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/analytics/overview")
	if overview.Clicks != 1 || overview.ImportedConversions != 1 || overview.CommissionCents != 12990 {
		t.Fatalf("analytics overview = %+v", overview)
	}

	topProducts := getJSON[[]topProductResponse](t, client, "/api/v1/workspaces/"+workspaceID+"/analytics/top-products")
	if len(topProducts) != 1 || topProducts[0].ProductID != product.ID || topProducts[0].Clicks != 1 || topProducts[0].ImportedConversions != 1 {
		t.Fatalf("top products = %+v", topProducts)
	}
}

func openIsolatedPool(t *testing.T, ctx context.Context, databaseURL string) *pgxpool.Pool {
	t.Helper()

	adminPool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open admin postgres pool: %v", err)
	}

	schema := fmt.Sprintf("test_%d", time.Now().UnixNano())
	if _, err := adminPool.Exec(ctx, "CREATE SCHEMA "+schema); err != nil {
		t.Fatalf("create test schema: %v", err)
	}
	t.Cleanup(func() {
		_, _ = adminPool.Exec(context.Background(), "DROP SCHEMA IF EXISTS "+schema+" CASCADE")
		adminPool.Close()
	})

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		t.Fatalf("parse postgres config: %v", err)
	}
	cfg.ConnConfig.RuntimeParams["search_path"] = schema

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		t.Fatalf("open isolated postgres pool: %v", err)
	}
	return pool
}

func runMigrations(t *testing.T, ctx context.Context, pool *pgxpool.Pool) {
	t.Helper()

	files, err := filepath.Glob("../../migrations/*.up.sql")
	if err != nil {
		t.Fatalf("list migrations: %v", err)
	}
	sort.Strings(files)
	for _, file := range files {
		sql, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("read migration %s: %v", file, err)
		}
		if _, err := pool.Exec(ctx, string(sql)); err != nil {
			t.Fatalf("apply migration %s: %v", filepath.Base(file), err)
		}
	}
}

type testClient struct {
	router  *gin.Engine
	cookies []*http.Cookie
}

func newTestClient(router *gin.Engine) *testClient {
	return &testClient{router: router}
}

func postJSON[T any](t *testing.T, client *testClient, path string, body string) T {
	t.Helper()

	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	return doJSON[T](t, client, req, http.StatusCreated)
}

func getJSON[T any](t *testing.T, client *testClient, path string) T {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, path, nil)
	return doJSON[T](t, client, req, http.StatusOK)
}

func patchJSON[T any](t *testing.T, client *testClient, path string, body string) T {
	t.Helper()

	req := httptest.NewRequest(http.MethodPatch, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	return doJSON[T](t, client, req, http.StatusOK)
}

func doJSON[T any](t *testing.T, client *testClient, req *http.Request, wantStatus int) T {
	t.Helper()

	for _, cookie := range client.cookies {
		req.AddCookie(cookie)
	}
	rec := httptest.NewRecorder()
	client.router.ServeHTTP(rec, req)
	if rec.Code != wantStatus {
		t.Fatalf("%s %s status = %d, want %d, body = %s", req.Method, req.URL.String(), rec.Code, wantStatus, rec.Body.String())
	}
	for _, cookie := range rec.Result().Cookies() {
		client.cookies = append(client.cookies, cookie)
	}

	var payload T
	if err := json.NewDecoder(strings.NewReader(rec.Body.String())).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v; body = %s", err, rec.Body.String())
	}
	return payload
}

type authResponse struct {
	Workspaces []membershipResponse `json:"workspaces"`
}

type membershipResponse struct {
	WorkspaceID string `json:"workspace_id"`
}

type workspaceProgramResponse struct {
	ID string `json:"id"`
}

type productResponse struct {
	ID string `json:"id"`
}

type campaignResponse struct {
	ID              string                   `json:"id"`
	ProductID       string                   `json:"product_id"`
	Status          string                   `json:"status"`
	ChannelPackages []channelPackageResponse `json:"channel_packages"`
}

type channelPackageResponse struct {
	ID string `json:"id"`
}

type publishingTaskResponse struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	CompletedAt string `json:"completed_at"`
}

type complianceCheckResponse struct {
	ID       string                      `json:"id"`
	Findings []complianceFindingResponse `json:"findings"`
}

type complianceFindingResponse struct {
	Code string `json:"code"`
}

type offerResponse struct {
	ID string `json:"id"`
}

type linkResponse struct {
	ID string `json:"id"`
}

type conversionImportResponse struct {
	ID     string                        `json:"id"`
	Source string                        `json:"source"`
	Rows   []conversionImportRowResponse `json:"rows"`
}

type conversionImportRowResponse struct {
	ID string `json:"id"`
}

type shortLinkResponse struct {
	Slug string `json:"slug"`
}

type clickMetricsResponse struct {
	GroupBy string            `json:"group_by"`
	Items   []clickMetricItem `json:"items"`
}

type clickMetricItem struct {
	GroupID string `json:"group_id"`
	Clicks  int64  `json:"clicks"`
}

type analyticsOverviewResponse struct {
	Clicks              int64 `json:"clicks"`
	ImportedConversions int64 `json:"imported_conversions"`
	CommissionCents     int64 `json:"commission_cents"`
}

type topProductResponse struct {
	ProductID           string `json:"product_id"`
	Clicks              int64  `json:"clicks"`
	ImportedConversions int64  `json:"imported_conversions"`
}
