package identity

import (
	"context"
	"errors"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
)

const sessionTTL = 14 * 24 * time.Hour

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrOAuthUnavailable   = errors.New("oauth provider is not configured")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (service *Service) Signup(ctx context.Context, input SignupInput) (AuthResponse, string, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return AuthResponse{}, "", err
	}

	passwordHash, err := hashPassword(input.Password)
	if err != nil {
		return AuthResponse{}, "", err
	}

	user, _, membership, err := service.repo.CreateUserWithWorkspace(ctx, input, passwordHash)
	if err != nil {
		return AuthResponse{}, "", err
	}

	session, token, err := service.createSession(ctx, user.ID)
	if err != nil {
		return AuthResponse{}, "", err
	}

	return AuthResponse{
		User:       user,
		Session:    SessionView{ExpiresAt: session.ExpiresAt},
		Workspaces: []Membership{membership},
	}, token, nil
}

func (service *Service) Login(ctx context.Context, input LoginInput) (AuthResponse, string, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return AuthResponse{}, "", err
	}

	user, passwordHash, err := service.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, common.ErrNotFound) {
			return AuthResponse{}, "", ErrInvalidCredentials
		}
		return AuthResponse{}, "", err
	}
	if user.Status != "active" {
		return AuthResponse{}, "", ErrInvalidCredentials
	}

	ok, err := verifyPassword(input.Password, passwordHash)
	if err != nil {
		return AuthResponse{}, "", err
	}
	if !ok {
		return AuthResponse{}, "", ErrInvalidCredentials
	}

	memberships, err := service.repo.ListMemberships(ctx, user.ID)
	if err != nil {
		return AuthResponse{}, "", err
	}
	session, token, err := service.createSession(ctx, user.ID)
	if err != nil {
		return AuthResponse{}, "", err
	}

	return AuthResponse{
		User:       user,
		Session:    SessionView{ExpiresAt: session.ExpiresAt},
		Workspaces: memberships,
	}, token, nil
}

func (service *Service) Logout(ctx context.Context, sessionToken string) error {
	if sessionToken == "" {
		return nil
	}
	return service.repo.RevokeSession(ctx, hashSessionToken(sessionToken))
}

func (service *Service) Authenticate(ctx context.Context, sessionToken string) (AuthenticatedUser, error) {
	if sessionToken == "" {
		return AuthenticatedUser{}, ErrUnauthorized
	}

	user, session, err := service.repo.GetUserBySession(ctx, hashSessionToken(sessionToken))
	if err != nil {
		if errors.Is(err, common.ErrNotFound) {
			return AuthenticatedUser{}, ErrUnauthorized
		}
		return AuthenticatedUser{}, err
	}
	memberships, err := service.repo.ListMemberships(ctx, user.ID)
	if err != nil {
		return AuthenticatedUser{}, err
	}
	return AuthenticatedUser{
		User:        user,
		SessionID:   session.ID,
		Memberships: memberships,
	}, nil
}

func (service *Service) Me(ctx context.Context, auth AuthenticatedUser) (MeResponse, error) {
	return MeResponse{
		User:       auth.User,
		Workspaces: auth.Memberships,
	}, nil
}

func (service *Service) AuthorizeWorkspace(ctx context.Context, userID string, workspaceID string, minimumRole string) (Membership, error) {
	membership, err := service.repo.GetMembership(ctx, userID, workspaceID)
	if err != nil {
		if errors.Is(err, common.ErrNotFound) {
			return Membership{}, ErrForbidden
		}
		return Membership{}, err
	}
	if !roleAllows(membership.Role, minimumRole) {
		return Membership{}, ErrForbidden
	}
	return membership, nil
}

func (service *Service) BeginOAuth(ctx context.Context, provider string, redirectURL string) error {
	if !isSupportedOAuthProvider(provider) {
		return common.NewValidationError("oauth provider is not supported")
	}
	return ErrOAuthUnavailable
}

func (service *Service) ListWorkspaces(ctx context.Context) ([]Workspace, error) {
	return service.repo.ListWorkspaces(ctx)
}

func (service *Service) ListMemberships(ctx context.Context, auth AuthenticatedUser) ([]Membership, error) {
	return service.repo.ListMemberships(ctx, auth.User.ID)
}

func (service *Service) CreateWorkspace(ctx context.Context, input CreateWorkspaceInput) (Workspace, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return Workspace{}, err
	}
	return service.repo.CreateWorkspace(ctx, input)
}

func (service *Service) CreateWorkspaceForUser(ctx context.Context, auth AuthenticatedUser, input CreateWorkspaceInput) (Workspace, Membership, error) {
	input.Normalize()
	if err := input.Validate(); err != nil {
		return Workspace{}, Membership{}, err
	}
	return service.repo.CreateWorkspaceForUser(ctx, auth.User.ID, input, RoleOwner)
}

func (service *Service) GetWorkspace(ctx context.Context, workspaceID string) (Workspace, error) {
	return service.repo.GetWorkspace(ctx, workspaceID)
}

func (service *Service) createSession(ctx context.Context, userID string) (Session, string, error) {
	token, sessionHash, err := newSessionToken()
	if err != nil {
		return Session{}, "", err
	}
	session, err := service.repo.CreateSession(ctx, userID, sessionHash, time.Now().UTC().Add(sessionTTL))
	if err != nil {
		return Session{}, "", err
	}
	return session, token, nil
}

func roleAllows(actual string, minimum string) bool {
	ranks := map[string]int{
		RoleMember: 1,
		RoleAdmin:  2,
		RoleOwner:  3,
	}
	return ranks[actual] >= ranks[minimum]
}

func isSupportedOAuthProvider(provider string) bool {
	switch provider {
	case OAuthProviderGoogle, OAuthProviderTikTok, OAuthProviderAmazon:
		return true
	default:
		return false
	}
}
