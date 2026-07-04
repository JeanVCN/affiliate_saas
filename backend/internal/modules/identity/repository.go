package identity

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	ListWorkspaces(ctx context.Context) ([]Workspace, error)
	CreateWorkspace(ctx context.Context, input CreateWorkspaceInput) (Workspace, error)
	CreateWorkspaceForUser(ctx context.Context, userID string, input CreateWorkspaceInput, role string) (Workspace, Membership, error)
	GetWorkspace(ctx context.Context, workspaceID string) (Workspace, error)
	CreateUserWithWorkspace(ctx context.Context, input SignupInput, passwordHash string) (User, Workspace, Membership, error)
	GetUserByEmail(ctx context.Context, email string) (User, string, error)
	GetUserBySession(ctx context.Context, sessionHash string) (User, Session, error)
	CreateSession(ctx context.Context, userID string, sessionHash string, expiresAt time.Time) (Session, error)
	RevokeSession(ctx context.Context, sessionHash string) error
	ListMemberships(ctx context.Context, userID string) ([]Membership, error)
	GetMembership(ctx context.Context, userID string, workspaceID string) (Membership, error)
	CreateOAuthState(ctx context.Context, input OAuthState) (OAuthState, error)
	GetOAuthIdentity(ctx context.Context, provider string, providerSubject string) (OAuthIdentity, error)
}

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}
