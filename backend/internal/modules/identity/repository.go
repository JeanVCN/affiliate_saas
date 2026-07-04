package identity

import (
	"context"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ListWorkspaces(ctx context.Context) ([]Workspace, error)
	CreateWorkspace(ctx context.Context, input CreateWorkspaceInput) (Workspace, error)
	GetWorkspace(ctx context.Context, workspaceID string) (Workspace, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) ListWorkspaces(ctx context.Context) ([]Workspace, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, name, slug, status, created_at, updated_at
		FROM workspaces
		WHERE archived_at IS NULL
		ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Workspace
	for rows.Next() {
		var item Workspace
		if err := rows.Scan(&item.ID, &item.Name, &item.Slug, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) CreateWorkspace(ctx context.Context, input CreateWorkspaceInput) (Workspace, error) {
	item := Workspace{ID: common.NewID("wks")}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO workspaces (id, name, slug)
		VALUES ($1, $2, $3)
		RETURNING id, name, slug, status, created_at, updated_at`,
		item.ID, input.Name, input.Slug,
	).Scan(&item.ID, &item.Name, &item.Slug, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) GetWorkspace(ctx context.Context, workspaceID string) (Workspace, error) {
	var item Workspace
	err := repo.db.QueryRow(ctx, `
		SELECT id, name, slug, status, created_at, updated_at
		FROM workspaces
		WHERE id = $1 AND archived_at IS NULL`, workspaceID,
	).Scan(&item.ID, &item.Name, &item.Slug, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}
