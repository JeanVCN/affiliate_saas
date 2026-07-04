package analytics

import (
	"net/http"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func RegisterRoutes(api gin.IRouter, service *Service) {
	handler := NewHandler(service)
	api.GET("/workspaces/:workspace_id/analytics/clicks", handler.ClickMetrics)
}

func (handler *Handler) ClickMetrics(c *gin.Context) {
	groupBy := c.DefaultQuery("group_by", "product")
	items, err := handler.service.ClickMetrics(c.Request.Context(), c.Param("workspace_id"), groupBy)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, gin.H{"group_by": groupBy, "items": items}, err, http.StatusOK)
}
