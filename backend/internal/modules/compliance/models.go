package compliance

import (
	"strings"
	"time"
)

type Check struct {
	ID          string    `json:"id"`
	WorkspaceID string    `json:"workspace_id"`
	CampaignID  string    `json:"campaign_id,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	Findings    []Finding `json:"findings"`
}

type Finding struct {
	ID                string    `json:"id"`
	WorkspaceID       string    `json:"workspace_id"`
	ComplianceCheckID string    `json:"compliance_check_id"`
	Severity          string    `json:"severity"`
	Code              string    `json:"code"`
	Message           string    `json:"message"`
	CreatedAt         time.Time `json:"created_at"`
}

type RunCampaignCheckInput struct {
	Channel string `json:"channel"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

type FindingInput struct {
	Severity string
	Code     string
	Message  string
}

type PolicyNote struct {
	Severity string
	Title    string
	Body     string
}

func (input *RunCampaignCheckInput) Normalize() {
	input.Channel = strings.ToLower(strings.TrimSpace(input.Channel))
	input.Title = strings.TrimSpace(input.Title)
	input.Body = strings.TrimSpace(input.Body)
}

func (input RunCampaignCheckInput) content() string {
	return strings.ToLower(strings.TrimSpace(input.Title + " " + input.Body))
}
