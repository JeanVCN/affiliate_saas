package marketplace

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) ListMarketplaces(ctx context.Context) ([]Marketplace, error) {
	return service.repo.ListMarketplaces(ctx)
}

func (service *Service) ListWorkspacePrograms(ctx context.Context, workspaceID string) ([]WorkspaceProgram, error) {
	return service.repo.ListWorkspacePrograms(ctx, workspaceID)
}

func (service *Service) EnableWorkspaceProgram(ctx context.Context, workspaceID string, input EnableWorkspaceProgramInput) (WorkspaceProgram, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return WorkspaceProgram{}, err
	}
	return service.repo.EnableWorkspaceProgram(ctx, workspaceID, input)
}
