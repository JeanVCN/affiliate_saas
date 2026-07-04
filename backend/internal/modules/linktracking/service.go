package linktracking

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

func (service *Service) CreateShortLink(ctx context.Context, workspaceID string, linkID string, input CreateShortLinkInput) (ShortLink, error) {
	input.Normalize()
	return service.repo.CreateShortLink(ctx, workspaceID, linkID, input)
}

func (service *Service) ResolveShortLink(ctx context.Context, slug string) (ResolvedShortLink, error) {
	return service.repo.ResolveShortLink(ctx, slug)
}

func (service *Service) RecordClick(ctx context.Context, input RecordClickInput) error {
	if input.ID == "" {
		input.ID = common.NewID("clk")
	}
	return service.repo.RecordClick(ctx, input)
}
