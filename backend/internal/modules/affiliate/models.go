package affiliate

import (
	"net/url"
	"strings"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

type Link struct {
	ID             string      `json:"id"`
	WorkspaceID    string      `json:"workspace_id"`
	ProductID      string      `json:"product_id"`
	OfferID        string      `json:"offer_id,omitempty"`
	DestinationURL string      `json:"destination_url"`
	Label          string      `json:"label,omitempty"`
	Status         string      `json:"status"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	ShortLinks     []ShortLink `json:"short_links,omitempty"`
}

type ShortLink struct {
	ID              string    `json:"id"`
	WorkspaceID     string    `json:"workspace_id"`
	AffiliateLinkID string    `json:"affiliate_link_id"`
	LinkVariantID   string    `json:"link_variant_id,omitempty"`
	Slug            string    `json:"slug"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CreateLinkInput struct {
	ProductID      string `json:"product_id"`
	OfferID        string `json:"offer_id"`
	DestinationURL string `json:"destination_url"`
	Label          string `json:"label"`
}

func (input *CreateLinkInput) Normalize() {
	input.ProductID = strings.TrimSpace(input.ProductID)
	input.OfferID = strings.TrimSpace(input.OfferID)
	input.DestinationURL = strings.TrimSpace(input.DestinationURL)
	input.Label = strings.TrimSpace(input.Label)
}

func (input CreateLinkInput) Validate() error {
	if input.ProductID == "" {
		return common.NewValidationError("product_id is required")
	}
	if input.DestinationURL == "" {
		return common.NewValidationError("destination_url is required")
	}
	parsed, err := url.ParseRequestURI(input.DestinationURL)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return common.NewValidationError("destination_url must be an absolute URL")
	}
	return nil
}
