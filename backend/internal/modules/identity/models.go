package identity

import (
	"strings"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

type Workspace struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateWorkspaceInput struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (input *CreateWorkspaceInput) Normalize() {
	input.Name = strings.TrimSpace(input.Name)
	input.Slug = common.NormalizeSlug(input.Slug)
	if input.Slug == "" {
		input.Slug = common.NormalizeSlug(input.Name)
	}
}

func (input CreateWorkspaceInput) Validate() error {
	if input.Name == "" {
		return common.NewValidationError("name is required")
	}
	if input.Slug == "" {
		return common.NewValidationError("slug is required")
	}
	return nil
}
