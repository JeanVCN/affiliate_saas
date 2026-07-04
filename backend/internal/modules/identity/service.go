package identity

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) ListWorkspaces(ctx context.Context) ([]Workspace, error) {
	return service.repo.ListWorkspaces(ctx)
}

func (service *Service) CreateWorkspace(ctx context.Context, input CreateWorkspaceInput) (Workspace, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return Workspace{}, err
	}
	return service.repo.CreateWorkspace(ctx, input)
}

func (service *Service) GetWorkspace(ctx context.Context, workspaceID string) (Workspace, error) {
	return service.repo.GetWorkspace(ctx, workspaceID)
}
