package identity

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
	api.GET("/workspaces", handler.ListWorkspaces)
	api.POST("/workspaces", handler.CreateWorkspace)
	api.GET("/workspaces/:workspace_id", handler.GetWorkspace)
}

func (handler *Handler) ListWorkspaces(c *gin.Context) {
	items, err := handler.service.ListWorkspaces(c.Request.Context())
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) CreateWorkspace(c *gin.Context) {
	var input CreateWorkspaceInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateWorkspace(c.Request.Context(), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) GetWorkspace(c *gin.Context) {
	item, err := handler.service.GetWorkspace(c.Request.Context(), c.Param("workspace_id"))
	common.Respond(c, item, err, http.StatusOK)
}
