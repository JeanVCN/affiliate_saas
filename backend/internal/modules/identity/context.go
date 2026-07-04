package identity

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SessionCookieName = "affiliate_session"
	authContextKey    = "identity.authenticated_user"
)

func SetAuthenticatedUser(c *gin.Context, auth AuthenticatedUser) {
	c.Set(authContextKey, auth)
}

func GetAuthenticatedUser(c *gin.Context) (AuthenticatedUser, bool) {
	value, exists := c.Get(authContextKey)
	if !exists {
		return AuthenticatedUser{}, false
	}
	auth, ok := value.(AuthenticatedUser)
	return auth, ok
}

func RequireAuthenticatedUser(c *gin.Context) (AuthenticatedUser, bool) {
	auth, ok := GetAuthenticatedUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "unauthorized", "message": "authentication required"}})
		return AuthenticatedUser{}, false
	}
	return auth, true
}
