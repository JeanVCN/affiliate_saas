package compliance

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
	api.POST("/workspaces/:workspace_id/campaigns/:campaign_id/compliance-checks", handler.RunCampaignCheck)
}

func RegisterWorkspaceRoutes(workspace gin.IRouter, service *Service) {
	handler := NewHandler(service)
	workspace.POST("/campaigns/:campaign_id/compliance-checks", handler.RunCampaignCheck)
}

func (handler *Handler) RunCampaignCheck(c *gin.Context) {
	var input RunCampaignCheckInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.RunCampaignCheck(c.Request.Context(), c.Param("workspace_id"), c.Param("campaign_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}
