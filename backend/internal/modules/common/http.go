package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, target any) bool {
	if err := c.ShouldBindJSON(target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "bad_request", "message": "invalid JSON body"}})
		return false
	}
	return true
}

func Respond(c *gin.Context, payload any, err error, successStatus int) {
	if err == nil {
		c.JSON(successStatus, payload)
		return
	}
	switch {
	case errors.Is(err, ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "not_found", "message": "resource not found"}})
	case errors.Is(err, ErrConflict):
		c.JSON(http.StatusConflict, gin.H{"error": gin.H{"code": "conflict", "message": "resource already exists"}})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "internal_error", "message": "internal server error"}})
	}
}

func RespondValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "validation_error", "message": err.Error()}})
}
