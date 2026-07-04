package linktracking

import (
	"context"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateShortLink(ctx context.Context, workspaceID string, linkID string, input CreateShortLinkInput) (ShortLink, error)
	ResolveShortLink(ctx context.Context, slug string) (ResolvedShortLink, error)
	RecordClick(ctx context.Context, input RecordClickInput) error
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) CreateShortLink(ctx context.Context, workspaceID string, linkID string, input CreateShortLinkInput) (ShortLink, error) {
	if input.Slug == "" {
		input.Slug = common.NewSlug()
	}
	item := ShortLink{ID: common.NewID("sln")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO short_links (id, workspace_id, affiliate_link_id, slug)
		SELECT $1, $2, al.id, $4
		FROM affiliate_links al
		WHERE al.id = $3 AND al.workspace_id = $2 AND al.archived_at IS NULL
		RETURNING id, workspace_id, affiliate_link_id, COALESCE(link_variant_id, ''),
		          slug, status, created_at, updated_at`,
		item.ID, workspaceID, linkID, input.Slug,
	).Scan(&item.ID, &item.WorkspaceID, &item.AffiliateLinkID, &item.LinkVariantID, &item.Slug, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) ResolveShortLink(ctx context.Context, slug string) (ResolvedShortLink, error) {
	var item ResolvedShortLink
	err := repo.db.QueryRow(ctx, `
		SELECT sl.id, sl.workspace_id, sl.affiliate_link_id, al.product_id, al.destination_url,
		       COALESCE(lv.utm_source, ''), COALESCE(lv.utm_medium, ''), COALESCE(lv.utm_campaign, '')
		FROM short_links sl
		JOIN affiliate_links al ON al.id = sl.affiliate_link_id
		LEFT JOIN link_variants lv ON lv.id = sl.link_variant_id
		WHERE sl.slug = $1
		  AND sl.status = 'active'
		  AND al.status = 'active'
		  AND sl.archived_at IS NULL
		  AND al.archived_at IS NULL`,
		slug,
	).Scan(&item.ShortLinkID, &item.WorkspaceID, &item.AffiliateLinkID, &item.ProductID, &item.DestinationURL, &item.UTMSource, &item.UTMMedium, &item.UTMCampaign)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) RecordClick(ctx context.Context, input RecordClickInput) error {
	_, err := repo.db.Exec(ctx, `
		INSERT INTO click_events (
			id, workspace_id, short_link_id, affiliate_link_id, product_id,
			referrer, user_agent, ip_hash, utm_source, utm_medium, utm_campaign
		)
		VALUES ($1, $2, $3, $4, $5, NULLIF($6, ''), NULLIF($7, ''), NULLIF($8, ''),
		        NULLIF($9, ''), NULLIF($10, ''), NULLIF($11, ''))`,
		input.ID, input.WorkspaceID, input.ShortLinkID, input.AffiliateLinkID, input.ProductID,
		input.Referrer, input.UserAgent, input.IPHash, input.UTMSource, input.UTMMedium, input.UTMCampaign)
	return common.NormalizePostgresErr(err)
}
