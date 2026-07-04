package product

import (
	"strings"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

type Product struct {
	ID          string    `json:"id"`
	WorkspaceID string    `json:"workspace_id"`
	Name        string    `json:"name"`
	Category    string    `json:"category,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Offer struct {
	ID                 string    `json:"id"`
	WorkspaceID        string    `json:"workspace_id"`
	ProductID          string    `json:"product_id"`
	WorkspaceProgramID string    `json:"workspace_program_id"`
	Title              string    `json:"title,omitempty"`
	PriceCents         *int      `json:"price_cents,omitempty"`
	Currency           string    `json:"currency,omitempty"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CreateProductInput struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type CreateOfferInput struct {
	WorkspaceProgramID string `json:"workspace_program_id"`
	Title              string `json:"title"`
	PriceCents         *int   `json:"price_cents"`
	Currency           string `json:"currency"`
}

func (input *CreateProductInput) Normalize() {
	input.Name = strings.TrimSpace(input.Name)
	input.Category = strings.TrimSpace(input.Category)
	input.Description = strings.TrimSpace(input.Description)
}

func (input CreateProductInput) Validate() error {
	if input.Name == "" {
		return common.NewValidationError("name is required")
	}
	return nil
}

func (input *CreateOfferInput) Normalize() {
	input.WorkspaceProgramID = strings.TrimSpace(input.WorkspaceProgramID)
	input.Title = strings.TrimSpace(input.Title)
	input.Currency = strings.ToUpper(strings.TrimSpace(input.Currency))
}

func (input CreateOfferInput) Validate() error {
	if input.WorkspaceProgramID == "" {
		return common.NewValidationError("workspace_program_id is required")
	}
	if input.PriceCents != nil && *input.PriceCents < 0 {
		return common.NewValidationError("price_cents must be greater than or equal to zero")
	}
	return nil
}
