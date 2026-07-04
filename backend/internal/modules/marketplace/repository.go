package marketplace

import (
	"context"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ListMarketplaces(ctx context.Context) ([]Marketplace, error)
	ListWorkspacePrograms(ctx context.Context, workspaceID string) ([]WorkspaceProgram, error)
	EnableWorkspaceProgram(ctx context.Context, workspaceID string, input EnableWorkspaceProgramInput) (WorkspaceProgram, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) ListMarketplaces(ctx context.Context) ([]Marketplace, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, name, slug, status, created_at, updated_at
		FROM marketplaces
		WHERE status = 'active'
		ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Marketplace
	for rows.Next() {
		var item Marketplace
		if err := rows.Scan(&item.ID, &item.Name, &item.Slug, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) ListWorkspacePrograms(ctx context.Context, workspaceID string) ([]WorkspaceProgram, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT wp.id, wp.workspace_id, wp.program_id, p.marketplace_id, m.name, m.slug,
		       p.name, p.slug, COALESCE(wp.external_account_label, ''), wp.status,
		       wp.created_at, wp.updated_at
		FROM workspace_programs wp
		JOIN programs p ON p.id = wp.program_id
		JOIN marketplaces m ON m.id = p.marketplace_id
		WHERE wp.workspace_id = $1 AND wp.archived_at IS NULL
		ORDER BY wp.created_at DESC`, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []WorkspaceProgram
	for rows.Next() {
		var item WorkspaceProgram
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.ProgramID, &item.MarketplaceID, &item.MarketplaceName, &item.MarketplaceSlug, &item.ProgramName, &item.ProgramSlug, &item.ExternalAccountLabel, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) EnableWorkspaceProgram(ctx context.Context, workspaceID string, input EnableWorkspaceProgramInput) (WorkspaceProgram, error) {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return WorkspaceProgram{}, err
	}
	defer tx.Rollback(ctx)

	marketplaceID := common.NewID("mkt")
	err = tx.QueryRow(ctx, `
		INSERT INTO marketplaces (id, name, slug)
		VALUES ($1, $2, $3)
		ON CONFLICT (slug) DO UPDATE SET name = EXCLUDED.name, updated_at = now()
		RETURNING id`,
		marketplaceID, input.MarketplaceName, input.MarketplaceSlug,
	).Scan(&marketplaceID)
	if err != nil {
		return WorkspaceProgram{}, common.NormalizePostgresErr(err)
	}

	programID := common.NewID("prg")
	err = tx.QueryRow(ctx, `
		INSERT INTO programs (id, marketplace_id, name, slug)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (marketplace_id, slug) DO UPDATE SET name = EXCLUDED.name, updated_at = now()
		RETURNING id`,
		programID, marketplaceID, input.ProgramName, input.ProgramSlug,
	).Scan(&programID)
	if err != nil {
		return WorkspaceProgram{}, common.NormalizePostgresErr(err)
	}

	workspaceProgramID := common.NewID("wpr")
	err = tx.QueryRow(ctx, `
		INSERT INTO workspace_programs (id, workspace_id, program_id, external_account_label)
		VALUES ($1, $2, $3, NULLIF($4, ''))
		ON CONFLICT (workspace_id, program_id)
		DO UPDATE SET external_account_label = EXCLUDED.external_account_label, status = 'active', archived_at = NULL, updated_at = now()
		RETURNING id`,
		workspaceProgramID, workspaceID, programID, input.ExternalAccountLabel,
	).Scan(&workspaceProgramID)
	if err != nil {
		return WorkspaceProgram{}, common.NormalizePostgresErr(err)
	}

	if err := tx.Commit(ctx); err != nil {
		return WorkspaceProgram{}, err
	}
	return repo.getWorkspaceProgram(ctx, workspaceID, workspaceProgramID)
}

func (repo *PostgresRepository) getWorkspaceProgram(ctx context.Context, workspaceID string, workspaceProgramID string) (WorkspaceProgram, error) {
	var item WorkspaceProgram
	err := repo.db.QueryRow(ctx, `
		SELECT wp.id, wp.workspace_id, wp.program_id, p.marketplace_id, m.name, m.slug,
		       p.name, p.slug, COALESCE(wp.external_account_label, ''), wp.status,
		       wp.created_at, wp.updated_at
		FROM workspace_programs wp
		JOIN programs p ON p.id = wp.program_id
		JOIN marketplaces m ON m.id = p.marketplace_id
		WHERE wp.workspace_id = $1 AND wp.id = $2 AND wp.archived_at IS NULL`,
		workspaceID, workspaceProgramID,
	).Scan(&item.ID, &item.WorkspaceID, &item.ProgramID, &item.MarketplaceID, &item.MarketplaceName, &item.MarketplaceSlug, &item.ProgramName, &item.ProgramSlug, &item.ExternalAccountLabel, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}
