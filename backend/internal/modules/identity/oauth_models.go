package identity

import "time"

const (
	OAuthProviderGoogle = "google"
	OAuthProviderTikTok = "tiktok"
	OAuthProviderAmazon = "amazon"
)

type OAuthIdentity struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Provider        string    `json:"provider"`
	ProviderSubject string    `json:"provider_subject"`
	Email           string    `json:"email,omitempty"`
	DisplayName     string    `json:"display_name,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type OAuthState struct {
	ID          string
	Provider    string
	StateHash   string
	RedirectURL string
	ExpiresAt   time.Time
	ConsumedAt  *time.Time
	CreatedAt   time.Time
}
