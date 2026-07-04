package identity

import (
	"net/mail"
	"strings"
	"time"
	"unicode"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

const (
	RoleOwner  = "owner"
	RoleAdmin  = "admin"
	RoleMember = "member"
)

type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Membership struct {
	ID          string    `json:"id"`
	WorkspaceID string    `json:"workspace_id"`
	UserID      string    `json:"user_id"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Session struct {
	ID        string
	UserID    string
	ExpiresAt time.Time
	RevokedAt *time.Time
	CreatedAt time.Time
}

type SignupInput struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	DisplayName   string `json:"display_name"`
	WorkspaceName string `json:"workspace_name"`
	WorkspaceSlug string `json:"workspace_slug"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User       User         `json:"user"`
	Session    SessionView  `json:"session"`
	Workspaces []Membership `json:"workspaces"`
}

type MeResponse struct {
	User       User         `json:"user"`
	Workspaces []Membership `json:"workspaces"`
}

type SessionView struct {
	ExpiresAt time.Time `json:"expires_at"`
}

type AuthenticatedUser struct {
	User        User
	SessionID   string
	Memberships []Membership
}

func (input *SignupInput) Normalize() {
	input.Email = normalizeEmail(input.Email)
	input.DisplayName = strings.TrimSpace(input.DisplayName)
	input.WorkspaceName = strings.TrimSpace(input.WorkspaceName)
	input.WorkspaceSlug = common.NormalizeSlug(input.WorkspaceSlug)
	if input.WorkspaceName == "" && input.DisplayName != "" {
		input.WorkspaceName = input.DisplayName + " Workspace"
	}
	if input.WorkspaceSlug == "" {
		input.WorkspaceSlug = common.NormalizeSlug(input.WorkspaceName)
	}
}

func (input SignupInput) Validate() error {
	if err := validateEmail(input.Email); err != nil {
		return err
	}
	if err := validatePassword(input.Password); err != nil {
		return err
	}
	if input.WorkspaceName == "" {
		return common.NewValidationError("workspace_name is required")
	}
	if input.WorkspaceSlug == "" {
		return common.NewValidationError("workspace_slug is required")
	}
	return nil
}

func (input *LoginInput) Normalize() {
	input.Email = normalizeEmail(input.Email)
}

func (input LoginInput) Validate() error {
	if err := validateEmail(input.Email); err != nil {
		return err
	}
	if input.Password == "" {
		return common.NewValidationError("password is required")
	}
	return nil
}

func normalizeEmail(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func validateEmail(value string) error {
	if value == "" {
		return common.NewValidationError("email is required")
	}
	parsed, err := mail.ParseAddress(value)
	if err != nil || parsed.Address != value {
		return common.NewValidationError("email must be valid")
	}
	return nil
}

func validatePassword(value string) error {
	if len(value) < 12 {
		return common.NewValidationError("password must be at least 12 characters")
	}

	var hasLetter, hasNumber bool
	for _, r := range value {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsNumber(r) {
			hasNumber = true
		}
	}
	if !hasLetter || !hasNumber {
		return common.NewValidationError("password must include letters and numbers")
	}
	return nil
}
