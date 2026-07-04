package analytics

import (
	"context"
	"io"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) ClickMetrics(ctx context.Context, workspaceID string, groupBy string) ([]ClickMetric, error) {
	if groupBy == "" {
		groupBy = "product"
	}
	if groupBy != "product" && groupBy != "link" {
		return nil, common.NewValidationError("group_by must be product or link")
	}
	return service.repo.ClickMetrics(ctx, workspaceID, groupBy)
}

func (service *Service) Overview(ctx context.Context, workspaceID string) (Overview, error) {
	return service.repo.Overview(ctx, workspaceID)
}

func (service *Service) TopProducts(ctx context.Context, workspaceID string) ([]TopProduct, error) {
	return service.repo.TopProducts(ctx, workspaceID)
}

func (service *Service) CreateConversionImport(ctx context.Context, workspaceID string, input CreateConversionImportInput) (ConversionImport, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return ConversionImport{}, err
	}
	return service.repo.CreateConversionImport(ctx, workspaceID, input)
}

func (service *Service) CreateConversionImportRow(ctx context.Context, workspaceID string, importID string, input CreateConversionImportRowInput) (ConversionImportRow, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return ConversionImportRow{}, err
	}
	return service.repo.CreateConversionImportRow(ctx, workspaceID, importID, input)
}

func (service *Service) CreateConversionImportCSVRows(ctx context.Context, workspaceID string, importID string, input CreateConversionImportCSVInput) ([]ConversionImportRow, error) {
	rows, err := ParseConversionCSV(input.CSV)
	if err != nil {
		return nil, err
	}
	return service.repo.CreateConversionImportRows(ctx, workspaceID, importID, rows)
}

func (service *Service) CreateConversionImportCSVRowsFromReader(ctx context.Context, workspaceID string, importID string, reader io.Reader) ([]ConversionImportRow, error) {
	rows, err := ParseConversionCSVReader(reader)
	if err != nil {
		return nil, err
	}
	return service.repo.CreateConversionImportRows(ctx, workspaceID, importID, rows)
}

func (service *Service) GetConversionImport(ctx context.Context, workspaceID string, importID string) (ConversionImport, error) {
	return service.repo.GetConversionImport(ctx, workspaceID, importID)
}

func (service *Service) ReconciliationSummary(ctx context.Context, workspaceID string, importID string) (ReconciliationSummary, error) {
	return service.repo.ReconciliationSummary(ctx, workspaceID, importID)
}

func (service *Service) UpdateConversionImportRow(ctx context.Context, workspaceID string, importID string, rowID string, input UpdateConversionImportRowInput) (ConversionImportRow, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return ConversionImportRow{}, err
	}
	return service.repo.UpdateConversionImportRow(ctx, workspaceID, importID, rowID, input)
}
