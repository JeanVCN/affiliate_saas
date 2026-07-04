package identity

import (
	"context"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

func (repo *PostgresRepository) ListWorkspaces(ctx context.Context) ([]Workspace, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT w.id, w.name, w.slug, w.status, w.created_at, w.updated_at
		FROM workspaces w
		WHERE w.archived_at IS NULL
		ORDER BY w.created_at DESC`)
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

func (repo *PostgresRepository) CreateWorkspaceForUser(ctx context.Context, userID string, input CreateWorkspaceInput, role string) (Workspace, Membership, error) {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return Workspace{}, Membership{}, err
	}
	defer tx.Rollback(ctx)

	workspace := Workspace{ID: common.NewID("wks")}
	err = tx.QueryRow(ctx, `
		INSERT INTO workspaces (id, name, slug)
		VALUES ($1, $2, $3)
		RETURNING id, name, slug, status, created_at, updated_at`,
		workspace.ID, input.Name, input.Slug,
	).Scan(&workspace.ID, &workspace.Name, &workspace.Slug, &workspace.Status, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		return Workspace{}, Membership{}, common.NormalizePostgresErr(err)
	}

	membership := Membership{ID: common.NewID("wmb")}
	err = tx.QueryRow(ctx, `
		INSERT INTO workspace_memberships (id, workspace_id, user_id, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, workspace_id, user_id, role, status, created_at, updated_at`,
		membership.ID, workspace.ID, userID, role,
	).Scan(&membership.ID, &membership.WorkspaceID, &membership.UserID, &membership.Role, &membership.Status, &membership.CreatedAt, &membership.UpdatedAt)
	if err != nil {
		return Workspace{}, Membership{}, common.NormalizePostgresErr(err)
	}

	if err := tx.Commit(ctx); err != nil {
		return Workspace{}, Membership{}, err
	}
	return workspace, membership, nil
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
