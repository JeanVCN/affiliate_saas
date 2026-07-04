package marketplace

import (
	"strings"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

type Marketplace struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkspaceProgram struct {
	ID                   string    `json:"id"`
	WorkspaceID          string    `json:"workspace_id"`
	ProgramID            string    `json:"program_id"`
	MarketplaceID        string    `json:"marketplace_id"`
	MarketplaceName      string    `json:"marketplace_name"`
	MarketplaceSlug      string    `json:"marketplace_slug"`
	ProgramName          string    `json:"program_name"`
	ProgramSlug          string    `json:"program_slug"`
	ExternalAccountLabel string    `json:"external_account_label,omitempty"`
	Status               string    `json:"status"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type EnableWorkspaceProgramInput struct {
	MarketplaceName      string `json:"marketplace_name"`
	MarketplaceSlug      string `json:"marketplace_slug"`
	ProgramName          string `json:"program_name"`
	ProgramSlug          string `json:"program_slug"`
	ExternalAccountLabel string `json:"external_account_label"`
}

func (input *EnableWorkspaceProgramInput) Normalize() {
	input.MarketplaceName = strings.TrimSpace(input.MarketplaceName)
	input.MarketplaceSlug = common.NormalizeSlug(input.MarketplaceSlug)
	if input.MarketplaceSlug == "" {
		input.MarketplaceSlug = common.NormalizeSlug(input.MarketplaceName)
	}
	input.ProgramName = strings.TrimSpace(input.ProgramName)
	input.ProgramSlug = common.NormalizeSlug(input.ProgramSlug)
	if input.ProgramSlug == "" {
		input.ProgramSlug = common.NormalizeSlug(input.ProgramName)
	}
	input.ExternalAccountLabel = strings.TrimSpace(input.ExternalAccountLabel)
}

func (input EnableWorkspaceProgramInput) Validate() error {
	if input.MarketplaceName == "" {
		return common.NewValidationError("marketplace_name is required")
	}
	if input.MarketplaceSlug == "" {
		return common.NewValidationError("marketplace_slug is required")
	}
	if input.ProgramName == "" {
		return common.NewValidationError("program_name is required")
	}
	if input.ProgramSlug == "" {
		return common.NewValidationError("program_slug is required")
	}
	return nil
}
