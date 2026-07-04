package affiliate

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
	api.GET("/workspaces/:workspace_id/links", handler.ListLinks)
	api.POST("/workspaces/:workspace_id/links", handler.CreateLink)
	api.GET("/workspaces/:workspace_id/links/:link_id", handler.GetLink)
}

func (handler *Handler) ListLinks(c *gin.Context) {
	items, err := handler.service.ListLinks(c.Request.Context(), c.Param("workspace_id"))
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) CreateLink(c *gin.Context) {
	var input CreateLinkInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateLink(c.Request.Context(), c.Param("workspace_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) GetLink(c *gin.Context) {
	item, err := handler.service.GetLink(c.Request.Context(), c.Param("workspace_id"), c.Param("link_id"))
	common.Respond(c, item, err, http.StatusOK)
}
