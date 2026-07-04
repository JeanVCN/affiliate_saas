package product

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) ListProducts(ctx context.Context, workspaceID string) ([]Product, error) {
	return service.repo.ListProducts(ctx, workspaceID)
}

func (service *Service) CreateProduct(ctx context.Context, workspaceID string, input CreateProductInput) (Product, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return Product{}, err
	}
	return service.repo.CreateProduct(ctx, workspaceID, input)
}

func (service *Service) GetProduct(ctx context.Context, workspaceID string, productID string) (Product, error) {
	return service.repo.GetProduct(ctx, workspaceID, productID)
}

func (service *Service) CreateOffer(ctx context.Context, workspaceID string, productID string, input CreateOfferInput) (Offer, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return Offer{}, err
	}
	return service.repo.CreateOffer(ctx, workspaceID, productID, input)
}
