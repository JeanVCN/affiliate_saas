package affiliate

import (
	"context"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ListLinks(ctx context.Context, workspaceID string) ([]Link, error)
	CreateLink(ctx context.Context, workspaceID string, input CreateLinkInput) (Link, error)
	GetLink(ctx context.Context, workspaceID string, linkID string) (Link, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) ListLinks(ctx context.Context, workspaceID string) ([]Link, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, workspace_id, product_id, COALESCE(offer_id, ''), destination_url,
		       COALESCE(label, ''), status, created_at, updated_at
		FROM affiliate_links
		WHERE workspace_id = $1 AND archived_at IS NULL
		ORDER BY created_at DESC`, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Link
	for rows.Next() {
		var item Link
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.ProductID, &item.OfferID, &item.DestinationURL, &item.Label, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) CreateLink(ctx context.Context, workspaceID string, input CreateLinkInput) (Link, error) {
	item := Link{ID: common.NewID("lnk")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO affiliate_links (id, workspace_id, product_id, offer_id, destination_url, label)
		SELECT $1, $2, p.id, NULLIF($4, ''), $5, NULLIF($6, '')
		FROM products p
		WHERE p.id = $3
		  AND p.workspace_id = $2
		  AND p.archived_at IS NULL
		  AND (
		    NULLIF($4, '') IS NULL
		    OR EXISTS (
		      SELECT 1
		      FROM offers o
		      WHERE o.id = $4
		        AND o.workspace_id = $2
		        AND o.product_id = p.id
		        AND o.archived_at IS NULL
		    )
		  )
		RETURNING id, workspace_id, product_id, COALESCE(offer_id, ''), destination_url,
		          COALESCE(label, ''), status, created_at, updated_at`,
		item.ID, workspaceID, input.ProductID, input.OfferID, input.DestinationURL, input.Label,
	).Scan(&item.ID, &item.WorkspaceID, &item.ProductID, &item.OfferID, &item.DestinationURL, &item.Label, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) GetLink(ctx context.Context, workspaceID string, linkID string) (Link, error) {
	var item Link
	err := repo.db.QueryRow(ctx, `
		SELECT id, workspace_id, product_id, COALESCE(offer_id, ''), destination_url,
		       COALESCE(label, ''), status, created_at, updated_at
		FROM affiliate_links
		WHERE workspace_id = $1 AND id = $2 AND archived_at IS NULL`,
		workspaceID, linkID,
	).Scan(&item.ID, &item.WorkspaceID, &item.ProductID, &item.OfferID, &item.DestinationURL, &item.Label, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return item, common.NormalizePostgresErr(err)
	}
	item.ShortLinks, err = repo.listShortLinks(ctx, workspaceID, linkID)
	return item, err
}

func (repo *PostgresRepository) listShortLinks(ctx context.Context, workspaceID string, linkID string) ([]ShortLink, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, workspace_id, affiliate_link_id, COALESCE(link_variant_id, ''),
		       slug, status, created_at, updated_at
		FROM short_links
		WHERE workspace_id = $1 AND affiliate_link_id = $2 AND archived_at IS NULL
		ORDER BY created_at DESC`, workspaceID, linkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ShortLink
	for rows.Next() {
		var item ShortLink
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.AffiliateLinkID, &item.LinkVariantID, &item.Slug, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}
