package marketplace

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
	api.GET("/marketplaces", handler.ListMarketplaces)
	api.GET("/workspaces/:workspace_id/programs", handler.ListWorkspacePrograms)
	api.POST("/workspaces/:workspace_id/programs", handler.EnableWorkspaceProgram)
}

func (handler *Handler) ListMarketplaces(c *gin.Context) {
	items, err := handler.service.ListMarketplaces(c.Request.Context())
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) ListWorkspacePrograms(c *gin.Context) {
	items, err := handler.service.ListWorkspacePrograms(c.Request.Context(), c.Param("workspace_id"))
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) EnableWorkspaceProgram(c *gin.Context) {
	var input EnableWorkspaceProgramInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.EnableWorkspaceProgram(c.Request.Context(), c.Param("workspace_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}
