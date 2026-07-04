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

	workspace := postJSON[workspaceResponse](t, router, "/api/v1/workspaces", `{"name":"Creator Lab"}`)
	program := postJSON[workspaceProgramResponse](t, router, "/api/v1/workspaces/"+workspace.ID+"/programs", `{
		"marketplace_name":"TikTok Shop",
		"program_name":"TikTok Shop Affiliate",
		"external_account_label":"main"
	}`)
	product := postJSON[productResponse](t, router, "/api/v1/workspaces/"+workspace.ID+"/products", `{
		"name":"Creator Camera",
		"category":"electronics"
	}`)
	offer := postJSON[offerResponse](t, router, "/api/v1/workspaces/"+workspace.ID+"/products/"+product.ID+"/offers", fmt.Sprintf(`{
		"workspace_program_id":%q,
		"title":"TikTok Shop offer",
		"price_cents":129900,
		"currency":"brl"
	}`, program.ID))
	link := postJSON[linkResponse](t, router, "/api/v1/workspaces/"+workspace.ID+"/links", fmt.Sprintf(`{
		"product_id":%q,
		"offer_id":%q,
		"destination_url":"https://example.com/products/creator-camera",
		"label":"TikTok bio"
	}`, product.ID, offer.ID))
	shortLink := postJSON[shortLinkResponse](t, router, "/api/v1/workspaces/"+workspace.ID+"/links/"+link.ID+"/short-links", `{"slug":"creator-camera"}`)

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

	clicksByProduct := getJSON[clickMetricsResponse](t, router, "/api/v1/workspaces/"+workspace.ID+"/analytics/clicks?group_by=product")
	if clicksByProduct.GroupBy != "product" || len(clicksByProduct.Items) != 1 {
		t.Fatalf("product metrics = %+v", clicksByProduct)
	}
	if clicksByProduct.Items[0].GroupID != product.ID || clicksByProduct.Items[0].Clicks != 1 {
		t.Fatalf("product metric item = %+v", clicksByProduct.Items[0])
	}

	clicksByLink := getJSON[clickMetricsResponse](t, router, "/api/v1/workspaces/"+workspace.ID+"/analytics/clicks?group_by=link")
	if clicksByLink.GroupBy != "link" || len(clicksByLink.Items) != 1 {
		t.Fatalf("link metrics = %+v", clicksByLink)
	}
	if clicksByLink.Items[0].GroupID != link.ID || clicksByLink.Items[0].Clicks != 1 {
		t.Fatalf("link metric item = %+v", clicksByLink.Items[0])
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

func postJSON[T any](t *testing.T, router *gin.Engine, path string, body string) T {
	t.Helper()

	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	return doJSON[T](t, router, req, http.StatusCreated)
}

func getJSON[T any](t *testing.T, router *gin.Engine, path string) T {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, path, nil)
	return doJSON[T](t, router, req, http.StatusOK)
}

func doJSON[T any](t *testing.T, router *gin.Engine, req *http.Request, wantStatus int) T {
	t.Helper()

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != wantStatus {
		t.Fatalf("%s %s status = %d, want %d, body = %s", req.Method, req.URL.String(), rec.Code, wantStatus, rec.Body.String())
	}

	var payload T
	if err := json.NewDecoder(strings.NewReader(rec.Body.String())).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v; body = %s", err, rec.Body.String())
	}
	return payload
}

type workspaceResponse struct {
	ID string `json:"id"`
}

type workspaceProgramResponse struct {
	ID string `json:"id"`
}

type productResponse struct {
	ID string `json:"id"`
}

type offerResponse struct {
	ID string `json:"id"`
}

type linkResponse struct {
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
