package linktracking

import (
	"strings"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

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

type CreateShortLinkInput struct {
	Slug string `json:"slug"`
}

type ResolvedShortLink struct {
	ShortLinkID     string
	WorkspaceID     string
	AffiliateLinkID string
	ProductID       string
	DestinationURL  string
	UTMSource       string
	UTMMedium       string
	UTMCampaign     string
}

type RecordClickInput struct {
	ID              string
	WorkspaceID     string
	ShortLinkID     string
	AffiliateLinkID string
	ProductID       string
	Referrer        string
	UserAgent       string
	IPHash          string
	UTMSource       string
	UTMMedium       string
	UTMCampaign     string
}

func (input *CreateShortLinkInput) Normalize() {
	input.Slug = common.NormalizeSlug(strings.TrimSpace(input.Slug))
}
