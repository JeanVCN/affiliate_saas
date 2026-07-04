package httpapi

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules"
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/identity"
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/linktracking"
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/product"
)

func TestCreateWorkspaceEndpoint(t *testing.T) {
	repo := &fakeIdentityRepository{}
	router := NewRouter(Dependencies{
		AppEnv: "test",
		Modules: modules.Dependencies{
			Identity: identity.NewService(repo),
		},
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/workspaces", bytes.NewBufferString(`{"name":"Creator Lab"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusCreated, rec.Body.String())
	}
	if repo.createdWorkspace.Slug != "creator-lab" {
		t.Fatalf("slug = %q, want creator-lab", repo.createdWorkspace.Slug)
	}
}

func TestCreateProductValidation(t *testing.T) {
	router := NewRouter(Dependencies{
		AppEnv: "test",
		Modules: modules.Dependencies{
			Product: product.NewService(&fakeProductRepository{}),
		},
	})

	req := httptest.NewRequest(http.MethodPost, "/api/v1/workspaces/wks_1/products", bytes.NewBufferString(`{"name":" "}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}

func TestRedirectRecordsClickAndContinues(t *testing.T) {
	repo := &fakeLinkTrackingRepository{
		resolved: linktracking.ResolvedShortLink{
			ShortLinkID:     "sln_1",
			WorkspaceID:     "wks_1",
			AffiliateLinkID: "lnk_1",
			ProductID:       "prd_1",
			DestinationURL:  "https://example.com/product",
			UTMSource:       "tiktok",
		},
		recordClickErr: errors.New("temporary analytics write failure"),
	}
	router := NewRouter(Dependencies{
		AppEnv: "test",
		Modules: modules.Dependencies{
			LinkTracking: linktracking.NewService(repo),
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/r/abc123", nil)
	req.Header.Set("User-Agent", "test-agent")
	req.Header.Set("Referer", "https://creator.example/post")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusFound)
	}
	if location := rec.Header().Get("Location"); location != "https://example.com/product" {
		t.Fatalf("Location = %q", location)
	}
	if repo.recordedClick.ShortLinkID != "sln_1" {
		t.Fatalf("recorded short link = %q, want sln_1", repo.recordedClick.ShortLinkID)
	}
	if repo.recordedClick.UTMSource != "tiktok" {
		t.Fatalf("utm source = %q, want tiktok", repo.recordedClick.UTMSource)
	}
}

type fakeIdentityRepository struct {
	createdWorkspace identity.CreateWorkspaceInput
}

func (repo *fakeIdentityRepository) ListWorkspaces(context.Context) ([]identity.Workspace, error) {
	return []identity.Workspace{}, nil
}

func (repo *fakeIdentityRepository) CreateWorkspace(_ context.Context, input identity.CreateWorkspaceInput) (identity.Workspace, error) {
	repo.createdWorkspace = input
	return identity.Workspace{
		ID:        "wks_1",
		Name:      input.Name,
		Slug:      input.Slug,
		Status:    "active",
		CreatedAt: time.Unix(1, 0).UTC(),
		UpdatedAt: time.Unix(1, 0).UTC(),
	}, nil
}

func (repo *fakeIdentityRepository) GetWorkspace(context.Context, string) (identity.Workspace, error) {
	return identity.Workspace{ID: "wks_1", Name: "Creator Lab", Slug: "creator-lab", Status: "active"}, nil
}

type fakeProductRepository struct{}

func (repo *fakeProductRepository) ListProducts(context.Context, string) ([]product.Product, error) {
	return []product.Product{}, nil
}

func (repo *fakeProductRepository) CreateProduct(_ context.Context, workspaceID string, input product.CreateProductInput) (product.Product, error) {
	return product.Product{ID: "prd_1", WorkspaceID: workspaceID, Name: input.Name, Status: "active"}, nil
}

func (repo *fakeProductRepository) GetProduct(context.Context, string, string) (product.Product, error) {
	return product.Product{ID: "prd_1", WorkspaceID: "wks_1", Name: "Camera", Status: "active"}, nil
}

func (repo *fakeProductRepository) CreateOffer(_ context.Context, workspaceID string, productID string, input product.CreateOfferInput) (product.Offer, error) {
	return product.Offer{ID: "off_1", WorkspaceID: workspaceID, ProductID: productID, WorkspaceProgramID: input.WorkspaceProgramID, Status: "active"}, nil
}

type fakeLinkTrackingRepository struct {
	resolved       linktracking.ResolvedShortLink
	recordedClick  linktracking.RecordClickInput
	recordClickErr error
}

func (repo *fakeLinkTrackingRepository) CreateShortLink(_ context.Context, workspaceID string, linkID string, input linktracking.CreateShortLinkInput) (linktracking.ShortLink, error) {
	if input.Slug == "" {
		input.Slug = "abc123"
	}
	return linktracking.ShortLink{ID: "sln_1", WorkspaceID: workspaceID, AffiliateLinkID: linkID, Slug: input.Slug, Status: "active"}, nil
}

func (repo *fakeLinkTrackingRepository) ResolveShortLink(context.Context, string) (linktracking.ResolvedShortLink, error) {
	return repo.resolved, nil
}

func (repo *fakeLinkTrackingRepository) RecordClick(_ context.Context, input linktracking.RecordClickInput) error {
	repo.recordedClick = input
	return repo.recordClickErr
}
