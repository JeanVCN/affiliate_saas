package campaign

import (
	"strings"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

const (
	StatusDraft     = "draft"
	StatusReady     = "ready"
	StatusPublished = "published"
	StatusArchived  = "archived"

	PackageStatusDraft    = "draft"
	PackageStatusReady    = "ready"
	PackageStatusArchived = "archived"

	TaskStatusTodo      = "todo"
	TaskStatusScheduled = "scheduled"
	TaskStatusDone      = "done"
	TaskStatusCanceled  = "canceled"
)

type Campaign struct {
	ID          string           `json:"id"`
	WorkspaceID string           `json:"workspace_id"`
	ProductID   string           `json:"product_id,omitempty"`
	Name        string           `json:"name"`
	Status      string           `json:"status"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Packages    []ChannelPackage `json:"channel_packages,omitempty"`
}

type ChannelPackage struct {
	ID          string    `json:"id"`
	WorkspaceID string    `json:"workspace_id"`
	CampaignID  string    `json:"campaign_id"`
	Channel     string    `json:"channel"`
	Title       string    `json:"title,omitempty"`
	Body        string    `json:"body,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PublishingTask struct {
	ID               string     `json:"id"`
	WorkspaceID      string     `json:"workspace_id"`
	CampaignID       string     `json:"campaign_id"`
	ChannelPackageID string     `json:"channel_package_id,omitempty"`
	Channel          string     `json:"channel"`
	Title            string     `json:"title"`
	Notes            string     `json:"notes,omitempty"`
	ScheduledFor     *time.Time `json:"scheduled_for,omitempty"`
	Status           string     `json:"status"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	CompletedAt      *time.Time `json:"completed_at,omitempty"`
}

type CreateCampaignInput struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
}

type UpdateCampaignInput struct {
	Name   *string `json:"name"`
	Status *string `json:"status"`
}

type CreateChannelPackageInput struct {
	Channel string `json:"channel"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type CreatePublishingTaskInput struct {
	ChannelPackageID string     `json:"channel_package_id"`
	Channel          string     `json:"channel"`
	Title            string     `json:"title"`
	Notes            string     `json:"notes"`
	ScheduledFor     *time.Time `json:"scheduled_for"`
}

type UpdatePublishingTaskInput struct {
	Status       *string    `json:"status"`
	Notes        *string    `json:"notes"`
	ScheduledFor *time.Time `json:"scheduled_for"`
}

func (input *CreateCampaignInput) Normalize() {
	input.ProductID = strings.TrimSpace(input.ProductID)
	input.Name = strings.TrimSpace(input.Name)
}

func (input CreateCampaignInput) Validate() error {
	if input.Name == "" {
		return common.NewValidationError("name is required")
	}
	return nil
}

func (input *UpdateCampaignInput) Normalize() {
	if input.Name != nil {
		value := strings.TrimSpace(*input.Name)
		input.Name = &value
	}
	if input.Status != nil {
		value := strings.ToLower(strings.TrimSpace(*input.Status))
		input.Status = &value
	}
}

func (input UpdateCampaignInput) Validate() error {
	if input.Name != nil && *input.Name == "" {
		return common.NewValidationError("name cannot be empty")
	}
	if input.Status != nil && !isCampaignStatus(*input.Status) {
		return common.NewValidationError("status must be draft, ready, published, or archived")
	}
	return nil
}

func (input *CreateChannelPackageInput) Normalize() {
	input.Channel = strings.ToLower(strings.TrimSpace(input.Channel))
	input.Title = strings.TrimSpace(input.Title)
	input.Body = strings.TrimSpace(input.Body)
}

func (input CreateChannelPackageInput) Validate() error {
	if input.Channel == "" {
		return common.NewValidationError("channel is required")
	}
	if input.Title == "" && input.Body == "" {
		return common.NewValidationError("title or body is required")
	}
	return nil
}

func (input *CreatePublishingTaskInput) Normalize() {
	input.ChannelPackageID = strings.TrimSpace(input.ChannelPackageID)
	input.Channel = strings.ToLower(strings.TrimSpace(input.Channel))
	input.Title = strings.TrimSpace(input.Title)
	input.Notes = strings.TrimSpace(input.Notes)
}

func (input CreatePublishingTaskInput) Validate() error {
	if input.Channel == "" {
		return common.NewValidationError("channel is required")
	}
	if input.Title == "" {
		return common.NewValidationError("title is required")
	}
	return nil
}

func (input *UpdatePublishingTaskInput) Normalize() {
	if input.Status != nil {
		value := strings.ToLower(strings.TrimSpace(*input.Status))
		input.Status = &value
	}
	if input.Notes != nil {
		value := strings.TrimSpace(*input.Notes)
		input.Notes = &value
	}
}

func (input UpdatePublishingTaskInput) Validate() error {
	if input.Status != nil && !isTaskStatus(*input.Status) {
		return common.NewValidationError("status must be todo, scheduled, done, or canceled")
	}
	return nil
}

func isCampaignStatus(value string) bool {
	switch value {
	case StatusDraft, StatusReady, StatusPublished, StatusArchived:
		return true
	default:
		return false
	}
}

func isTaskStatus(value string) bool {
	switch value {
	case TaskStatusTodo, TaskStatusScheduled, TaskStatusDone, TaskStatusCanceled:
		return true
	default:
		return false
	}
}
