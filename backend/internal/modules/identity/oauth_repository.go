package identity

import (
	"context"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

func (repo *PostgresRepository) CreateOAuthState(ctx context.Context, input OAuthState) (OAuthState, error) {
	item := input
	if item.ID == "" {
		item.ID = common.NewID("ost")
	}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO oauth_states (id, provider, state_hash, redirect_url, expires_at)
		VALUES ($1, $2, $3, NULLIF($4, ''), $5)
		RETURNING id, provider, state_hash, COALESCE(redirect_url, ''), expires_at, consumed_at, created_at`,
		item.ID, item.Provider, item.StateHash, item.RedirectURL, item.ExpiresAt,
	).Scan(&item.ID, &item.Provider, &item.StateHash, &item.RedirectURL, &item.ExpiresAt, &item.ConsumedAt, &item.CreatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) GetOAuthIdentity(ctx context.Context, provider string, providerSubject string) (OAuthIdentity, error) {
	var item OAuthIdentity
	err := repo.db.QueryRow(ctx, `
		SELECT id, user_id, provider, provider_subject, COALESCE(email, ''),
		       COALESCE(display_name, ''), created_at, updated_at
		FROM oauth_identities
		WHERE provider = $1 AND provider_subject = $2`,
		provider, providerSubject,
	).Scan(&item.ID, &item.UserID, &item.Provider, &item.ProviderSubject, &item.Email, &item.DisplayName, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}
