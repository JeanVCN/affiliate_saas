package analytics

import (
	"context"

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
