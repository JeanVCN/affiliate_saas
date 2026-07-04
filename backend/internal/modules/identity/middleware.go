package identity

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(service *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(SessionCookieName)
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "unauthorized", "message": "authentication required"}})
			c.Abort()
			return
		}

		auth, err := service.Authenticate(c.Request.Context(), token)
		if err != nil {
			if errors.Is(err, ErrUnauthorized) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "unauthorized", "message": "authentication required"}})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "internal_error", "message": "internal server error"}})
			c.Abort()
			return
		}

		SetAuthenticatedUser(c, auth)
		c.Next()
	}
}

func WorkspaceRoleMiddleware(service *Service, minimumRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, ok := RequireAuthenticatedUser(c)
		if !ok {
			c.Abort()
			return
		}
		workspaceID := c.Param("workspace_id")
		if workspaceID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "validation_error", "message": "workspace_id is required"}})
			c.Abort()
			return
		}
		if _, err := service.AuthorizeWorkspace(c.Request.Context(), auth.User.ID, workspaceID, minimumRole); err != nil {
			if errors.Is(err, ErrForbidden) {
				c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"code": "forbidden", "message": "workspace access denied"}})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "internal_error", "message": "internal server error"}})
			c.Abort()
			return
		}
		c.Next()
	}
}
