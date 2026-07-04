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

type ProgramPolicyNote struct {
	ID        string    `json:"id"`
	ProgramID string    `json:"program_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Severity  string    `json:"severity"`
	SourceURL string    `json:"source_url,omitempty"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EnableWorkspaceProgramInput struct {
	MarketplaceName      string `json:"marketplace_name"`
	MarketplaceSlug      string `json:"marketplace_slug"`
	ProgramName          string `json:"program_name"`
	ProgramSlug          string `json:"program_slug"`
	ExternalAccountLabel string `json:"external_account_label"`
}

type CreateProgramPolicyNoteInput struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Severity  string `json:"severity"`
	SourceURL string `json:"source_url"`
}

type UpdateProgramPolicyNoteInput struct {
	Title     *string `json:"title"`
	Body      *string `json:"body"`
	Severity  *string `json:"severity"`
	SourceURL *string `json:"source_url"`
	Status    *string `json:"status"`
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

func (input *CreateProgramPolicyNoteInput) Normalize() {
	input.Title = strings.TrimSpace(input.Title)
	input.Body = strings.TrimSpace(input.Body)
	input.Severity = strings.ToLower(strings.TrimSpace(input.Severity))
	if input.Severity == "" {
		input.Severity = "info"
	}
	input.SourceURL = strings.TrimSpace(input.SourceURL)
}

func (input CreateProgramPolicyNoteInput) Validate() error {
	if input.Title == "" {
		return common.NewValidationError("title is required")
	}
	if input.Body == "" {
		return common.NewValidationError("body is required")
	}
	if input.Severity != "info" && input.Severity != "warning" && input.Severity != "blocker" {
		return common.NewValidationError("severity must be info, warning, or blocker")
	}
	return nil
}

func (input *UpdateProgramPolicyNoteInput) Normalize() {
	trimStringPtr(&input.Title)
	trimStringPtr(&input.Body)
	trimStringPtr(&input.SourceURL)
	if input.Severity != nil {
		value := strings.ToLower(strings.TrimSpace(*input.Severity))
		input.Severity = &value
	}
	if input.Status != nil {
		value := strings.ToLower(strings.TrimSpace(*input.Status))
		input.Status = &value
	}
}

func (input UpdateProgramPolicyNoteInput) Validate() error {
	if input.Title != nil && *input.Title == "" {
		return common.NewValidationError("title cannot be empty")
	}
	if input.Body != nil && *input.Body == "" {
		return common.NewValidationError("body cannot be empty")
	}
	if input.Severity != nil && !isPolicySeverity(*input.Severity) {
		return common.NewValidationError("severity must be info, warning, or blocker")
	}
	if input.Status != nil && *input.Status != "active" && *input.Status != "archived" {
		return common.NewValidationError("status must be active or archived")
	}
	return nil
}

func isPolicySeverity(value string) bool {
	return value == "info" || value == "warning" || value == "blocker"
}

func trimStringPtr(value **string) {
	if *value == nil {
		return
	}
	trimmed := strings.TrimSpace(**value)
	*value = &trimmed
}
