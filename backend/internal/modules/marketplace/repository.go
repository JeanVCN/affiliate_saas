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
	CreateProgramPolicyNote(ctx context.Context, workspaceID string, programID string, input CreateProgramPolicyNoteInput) (ProgramPolicyNote, error)
	ListProgramPolicyNotes(ctx context.Context, workspaceID string, programID string) ([]ProgramPolicyNote, error)
	UpdateProgramPolicyNote(ctx context.Context, workspaceID string, programID string, noteID string, input UpdateProgramPolicyNoteInput) (ProgramPolicyNote, error)
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

func (repo *PostgresRepository) CreateProgramPolicyNote(ctx context.Context, workspaceID string, programID string, input CreateProgramPolicyNoteInput) (ProgramPolicyNote, error) {
	item := ProgramPolicyNote{ID: common.NewID("ppn")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO program_policy_notes (id, program_id, title, body, severity, source_url)
		SELECT $1, wp.program_id, $4, $5, $6, NULLIF($7, '')
		FROM workspace_programs wp
		WHERE wp.workspace_id = $2 AND wp.program_id = $3 AND wp.archived_at IS NULL
		RETURNING id, program_id, title, body, severity, COALESCE(source_url, ''), status, created_at, updated_at`,
		item.ID, workspaceID, programID, input.Title, input.Body, input.Severity, input.SourceURL,
	).Scan(&item.ID, &item.ProgramID, &item.Title, &item.Body, &item.Severity, &item.SourceURL, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) ListProgramPolicyNotes(ctx context.Context, workspaceID string, programID string) ([]ProgramPolicyNote, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT ppn.id, ppn.program_id, ppn.title, ppn.body, ppn.severity,
		       COALESCE(ppn.source_url, ''), ppn.status, ppn.created_at, ppn.updated_at
		FROM program_policy_notes ppn
		JOIN workspace_programs wp ON wp.program_id = ppn.program_id
		WHERE wp.workspace_id = $1 AND wp.program_id = $2 AND wp.archived_at IS NULL
		ORDER BY ppn.updated_at DESC`, workspaceID, programID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ProgramPolicyNote
	for rows.Next() {
		var item ProgramPolicyNote
		if err := rows.Scan(&item.ID, &item.ProgramID, &item.Title, &item.Body, &item.Severity, &item.SourceURL, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) UpdateProgramPolicyNote(ctx context.Context, workspaceID string, programID string, noteID string, input UpdateProgramPolicyNoteInput) (ProgramPolicyNote, error) {
	var item ProgramPolicyNote
	err := repo.db.QueryRow(ctx, `
		UPDATE program_policy_notes ppn
		SET title = COALESCE($4, ppn.title),
		    body = COALESCE($5, ppn.body),
		    severity = COALESCE($6, ppn.severity),
		    source_url = COALESCE(NULLIF($7, ''), ppn.source_url),
		    status = COALESCE($8, ppn.status),
		    reviewed_at = now(),
		    updated_at = now()
		FROM workspace_programs wp
		WHERE ppn.id = $3 AND ppn.program_id = $2
		  AND wp.workspace_id = $1 AND wp.program_id = ppn.program_id AND wp.archived_at IS NULL
		RETURNING ppn.id, ppn.program_id, ppn.title, ppn.body, ppn.severity,
		          COALESCE(ppn.source_url, ''), ppn.status, ppn.created_at, ppn.updated_at`,
		workspaceID, programID, noteID, input.Title, input.Body, input.Severity, input.SourceURL, input.Status,
	).Scan(&item.ID, &item.ProgramID, &item.Title, &item.Body, &item.Severity, &item.SourceURL, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
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
