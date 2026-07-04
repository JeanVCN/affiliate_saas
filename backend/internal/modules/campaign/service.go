package campaign

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) ListCampaigns(ctx context.Context, workspaceID string) ([]Campaign, error) {
	return service.repo.ListCampaigns(ctx, workspaceID)
}

func (service *Service) CreateCampaign(ctx context.Context, workspaceID string, input CreateCampaignInput) (Campaign, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return Campaign{}, err
	}
	return service.repo.CreateCampaign(ctx, workspaceID, input)
}

func (service *Service) GetCampaign(ctx context.Context, workspaceID string, campaignID string) (Campaign, error) {
	return service.repo.GetCampaign(ctx, workspaceID, campaignID)
}

func (service *Service) UpdateCampaign(ctx context.Context, workspaceID string, campaignID string, input UpdateCampaignInput) (Campaign, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return Campaign{}, err
	}
	return service.repo.UpdateCampaign(ctx, workspaceID, campaignID, input)
}

func (service *Service) CreateChannelPackage(ctx context.Context, workspaceID string, campaignID string, input CreateChannelPackageInput) (ChannelPackage, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return ChannelPackage{}, err
	}
	return service.repo.CreateChannelPackage(ctx, workspaceID, campaignID, input)
}

func (service *Service) ListPublishingTasks(ctx context.Context, workspaceID string, campaignID string) ([]PublishingTask, error) {
	return service.repo.ListPublishingTasks(ctx, workspaceID, campaignID)
}

func (service *Service) CreatePublishingTask(ctx context.Context, workspaceID string, campaignID string, input CreatePublishingTaskInput) (PublishingTask, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return PublishingTask{}, err
	}
	return service.repo.CreatePublishingTask(ctx, workspaceID, campaignID, input)
}

func (service *Service) UpdatePublishingTask(ctx context.Context, workspaceID string, campaignID string, taskID string, input UpdatePublishingTaskInput) (PublishingTask, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return PublishingTask{}, err
	}
	return service.repo.UpdatePublishingTask(ctx, workspaceID, campaignID, taskID, input)
}
