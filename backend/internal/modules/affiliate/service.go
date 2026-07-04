package affiliate

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) ListLinks(ctx context.Context, workspaceID string) ([]Link, error) {
	return service.repo.ListLinks(ctx, workspaceID)
}

func (service *Service) CreateLink(ctx context.Context, workspaceID string, input CreateLinkInput) (Link, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return Link{}, err
	}
	return service.repo.CreateLink(ctx, workspaceID, input)
}

func (service *Service) GetLink(ctx context.Context, workspaceID string, linkID string) (Link, error) {
	return service.repo.GetLink(ctx, workspaceID, linkID)
}
