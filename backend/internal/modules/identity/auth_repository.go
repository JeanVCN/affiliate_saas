package identity

import (
	"context"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

func (repo *PostgresRepository) CreateUserWithWorkspace(ctx context.Context, input SignupInput, passwordHash string) (User, Workspace, Membership, error) {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return User{}, Workspace{}, Membership{}, err
	}
	defer tx.Rollback(ctx)

	user := User{ID: common.NewID("usr")}
	err = tx.QueryRow(ctx, `
		INSERT INTO users (id, email, password_hash, display_name)
		VALUES ($1, $2, $3, NULLIF($4, ''))
		RETURNING id, email, COALESCE(display_name, ''), status, created_at, updated_at`,
		user.ID, input.Email, passwordHash, input.DisplayName,
	).Scan(&user.ID, &user.Email, &user.DisplayName, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return User{}, Workspace{}, Membership{}, common.NormalizePostgresErr(err)
	}

	workspace := Workspace{ID: common.NewID("wks")}
	err = tx.QueryRow(ctx, `
		INSERT INTO workspaces (id, name, slug)
		VALUES ($1, $2, $3)
		RETURNING id, name, slug, status, created_at, updated_at`,
		workspace.ID, input.WorkspaceName, input.WorkspaceSlug,
	).Scan(&workspace.ID, &workspace.Name, &workspace.Slug, &workspace.Status, &workspace.CreatedAt, &workspace.UpdatedAt)
	if err != nil {
		return User{}, Workspace{}, Membership{}, common.NormalizePostgresErr(err)
	}

	membership := Membership{ID: common.NewID("wmb")}
	err = tx.QueryRow(ctx, `
		INSERT INTO workspace_memberships (id, workspace_id, user_id, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, workspace_id, user_id, role, status, created_at, updated_at`,
		membership.ID, workspace.ID, user.ID, RoleOwner,
	).Scan(&membership.ID, &membership.WorkspaceID, &membership.UserID, &membership.Role, &membership.Status, &membership.CreatedAt, &membership.UpdatedAt)
	if err != nil {
		return User{}, Workspace{}, Membership{}, common.NormalizePostgresErr(err)
	}

	if err := tx.Commit(ctx); err != nil {
		return User{}, Workspace{}, Membership{}, err
	}
	return user, workspace, membership, nil
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (User, string, error) {
	var user User
	var passwordHash string
	err := repo.db.QueryRow(ctx, `
		SELECT id, email, password_hash, COALESCE(display_name, ''), status, created_at, updated_at
		FROM users
		WHERE lower(email) = lower($1) AND archived_at IS NULL`,
		email,
	).Scan(&user.ID, &user.Email, &passwordHash, &user.DisplayName, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	return user, passwordHash, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) GetUserBySession(ctx context.Context, sessionHash string) (User, Session, error) {
	var user User
	var session Session
	err := repo.db.QueryRow(ctx, `
		SELECT u.id, u.email, COALESCE(u.display_name, ''), u.status, u.created_at, u.updated_at,
		       s.id, s.user_id, s.expires_at, s.revoked_at, s.created_at
		FROM sessions s
		JOIN users u ON u.id = s.user_id
		WHERE s.id = $1
		  AND s.revoked_at IS NULL
		  AND s.expires_at > now()
		  AND u.status = 'active'
		  AND u.archived_at IS NULL`,
		sessionHash,
	).Scan(&user.ID, &user.Email, &user.DisplayName, &user.Status, &user.CreatedAt, &user.UpdatedAt, &session.ID, &session.UserID, &session.ExpiresAt, &session.RevokedAt, &session.CreatedAt)
	return user, session, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) CreateSession(ctx context.Context, userID string, sessionHash string, expiresAt time.Time) (Session, error) {
	session := Session{ID: sessionHash}
	err := repo.db.QueryRow(ctx, `
		INSERT INTO sessions (id, user_id, expires_at)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, expires_at, revoked_at, created_at`,
		sessionHash, userID, expiresAt,
	).Scan(&session.ID, &session.UserID, &session.ExpiresAt, &session.RevokedAt, &session.CreatedAt)
	return session, common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) RevokeSession(ctx context.Context, sessionHash string) error {
	_, err := repo.db.Exec(ctx, `
		UPDATE sessions
		SET revoked_at = now()
		WHERE id = $1 AND revoked_at IS NULL`,
		sessionHash,
	)
	return common.NormalizePostgresErr(err)
}

func (repo *PostgresRepository) ListMemberships(ctx context.Context, userID string) ([]Membership, error) {
	rows, err := repo.db.Query(ctx, `
		SELECT id, workspace_id, user_id, role, status, created_at, updated_at
		FROM workspace_memberships
		WHERE user_id = $1 AND status = 'active'
		ORDER BY created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Membership
	for rows.Next() {
		var item Membership
		if err := rows.Scan(&item.ID, &item.WorkspaceID, &item.UserID, &item.Role, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (repo *PostgresRepository) GetMembership(ctx context.Context, userID string, workspaceID string) (Membership, error) {
	var item Membership
	err := repo.db.QueryRow(ctx, `
		SELECT id, workspace_id, user_id, role, status, created_at, updated_at
		FROM workspace_memberships
		WHERE user_id = $1 AND workspace_id = $2 AND status = 'active'`,
		userID, workspaceID,
	).Scan(&item.ID, &item.WorkspaceID, &item.UserID, &item.Role, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, common.NormalizePostgresErr(err)
}
