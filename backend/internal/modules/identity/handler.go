package identity

import (
	"errors"
	"net/http"
	"time"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
	appEnv  string
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func RegisterRoutes(api gin.IRouter, service *Service) {
	handler := NewHandler(service)
	api.GET("/workspaces", handler.ListWorkspaces)
	api.POST("/workspaces", handler.CreateWorkspace)
	api.GET("/workspaces/:workspace_id", handler.GetWorkspace)
}

func RegisterAuthRoutes(api gin.IRouter, protected gin.IRouter, service *Service, appEnv string) {
	handler := NewHandler(service)
	handler.appEnv = appEnv
	api.POST("/auth/signup", handler.Signup)
	api.POST("/auth/login", handler.Login)
	api.GET("/auth/oauth/:provider/start", handler.BeginOAuth)
	protected.POST("/auth/logout", handler.Logout)
	protected.GET("/auth/me", handler.Me)
}

func (handler *Handler) Signup(c *gin.Context) {
	var input SignupInput
	if !common.BindJSON(c, &input) {
		return
	}
	payload, token, err := handler.service.Signup(c.Request.Context(), input)
	if handler.respondAuthError(c, err) {
		return
	}
	setSessionCookie(c, token, payload.Session.ExpiresAt, handler.appEnv)
	common.Respond(c, payload, err, http.StatusCreated)
}

func (handler *Handler) Login(c *gin.Context) {
	var input LoginInput
	if !common.BindJSON(c, &input) {
		return
	}
	payload, token, err := handler.service.Login(c.Request.Context(), input)
	if handler.respondAuthError(c, err) {
		return
	}
	setSessionCookie(c, token, payload.Session.ExpiresAt, handler.appEnv)
	common.Respond(c, payload, err, http.StatusOK)
}

func (handler *Handler) Logout(c *gin.Context) {
	token, _ := c.Cookie(SessionCookieName)
	err := handler.service.Logout(c.Request.Context(), token)
	clearSessionCookie(c, handler.appEnv)
	common.Respond(c, gin.H{"status": "ok"}, err, http.StatusOK)
}

func (handler *Handler) Me(c *gin.Context) {
	auth, ok := RequireAuthenticatedUser(c)
	if !ok {
		return
	}
	payload, err := handler.service.Me(c.Request.Context(), auth)
	common.Respond(c, payload, err, http.StatusOK)
}

func (handler *Handler) BeginOAuth(c *gin.Context) {
	err := handler.service.BeginOAuth(c.Request.Context(), c.Param("provider"), c.Query("redirect_url"))
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	if errors.Is(err, ErrOAuthUnavailable) {
		c.JSON(http.StatusNotImplemented, gin.H{"error": gin.H{"code": "oauth_unavailable", "message": "oauth provider is not configured"}})
		return
	}
	common.Respond(c, gin.H{"status": "ok"}, err, http.StatusOK)
}

func (handler *Handler) ListWorkspaces(c *gin.Context) {
	auth, ok := RequireAuthenticatedUser(c)
	if !ok {
		return
	}
	items, err := handler.service.ListMemberships(c.Request.Context(), auth)
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) CreateWorkspace(c *gin.Context) {
	var input CreateWorkspaceInput
	if !common.BindJSON(c, &input) {
		return
	}
	auth, ok := RequireAuthenticatedUser(c)
	if !ok {
		return
	}
	item, _, err := handler.service.CreateWorkspaceForUser(c.Request.Context(), auth, input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) GetWorkspace(c *gin.Context) {
	auth, ok := RequireAuthenticatedUser(c)
	if !ok {
		return
	}
	if _, err := handler.service.AuthorizeWorkspace(c.Request.Context(), auth.User.ID, c.Param("workspace_id"), RoleMember); err != nil {
		if errors.Is(err, ErrForbidden) {
			c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "forbidden", "message": "workspace access denied"}})
			return
		}
		common.Respond(c, nil, err, http.StatusOK)
		return
	}
	item, err := handler.service.GetWorkspace(c.Request.Context(), c.Param("workspace_id"))
	common.Respond(c, item, err, http.StatusOK)
}

func (handler *Handler) respondAuthError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return true
	}
	if errors.Is(err, ErrInvalidCredentials) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "invalid_credentials", "message": "invalid email or password"}})
		return true
	}
	common.Respond(c, nil, err, http.StatusOK)
	return true
}

func setSessionCookie(c *gin.Context, token string, expiresAt time.Time, appEnv string) {
	maxAge := int(time.Until(expiresAt).Seconds())
	if maxAge < 0 {
		maxAge = 0
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(SessionCookieName, token, maxAge, "/", "", appEnv != "development" && appEnv != "test", true)
}

func clearSessionCookie(c *gin.Context, appEnv string) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(SessionCookieName, "", -1, "/", "", appEnv != "development" && appEnv != "test", true)
}
