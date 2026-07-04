package campaign

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
	api.GET("/workspaces/:workspace_id/campaigns", handler.ListCampaigns)
	api.POST("/workspaces/:workspace_id/campaigns", handler.CreateCampaign)
	api.GET("/workspaces/:workspace_id/campaigns/:campaign_id", handler.GetCampaign)
	api.PATCH("/workspaces/:workspace_id/campaigns/:campaign_id", handler.UpdateCampaign)
	api.POST("/workspaces/:workspace_id/campaigns/:campaign_id/channel-packages", handler.CreateChannelPackage)
	api.GET("/workspaces/:workspace_id/campaigns/:campaign_id/publishing-tasks", handler.ListPublishingTasks)
	api.POST("/workspaces/:workspace_id/campaigns/:campaign_id/publishing-tasks", handler.CreatePublishingTask)
	api.PATCH("/workspaces/:workspace_id/campaigns/:campaign_id/publishing-tasks/:task_id", handler.UpdatePublishingTask)
}

func RegisterWorkspaceRoutes(workspace gin.IRouter, service *Service) {
	handler := NewHandler(service)
	workspace.GET("/campaigns", handler.ListCampaigns)
	workspace.POST("/campaigns", handler.CreateCampaign)
	workspace.GET("/campaigns/:campaign_id", handler.GetCampaign)
	workspace.PATCH("/campaigns/:campaign_id", handler.UpdateCampaign)
	workspace.POST("/campaigns/:campaign_id/channel-packages", handler.CreateChannelPackage)
	workspace.GET("/campaigns/:campaign_id/publishing-tasks", handler.ListPublishingTasks)
	workspace.POST("/campaigns/:campaign_id/publishing-tasks", handler.CreatePublishingTask)
	workspace.PATCH("/campaigns/:campaign_id/publishing-tasks/:task_id", handler.UpdatePublishingTask)
}

func (handler *Handler) ListCampaigns(c *gin.Context) {
	items, err := handler.service.ListCampaigns(c.Request.Context(), c.Param("workspace_id"))
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) CreateCampaign(c *gin.Context) {
	var input CreateCampaignInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateCampaign(c.Request.Context(), c.Param("workspace_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) GetCampaign(c *gin.Context) {
	item, err := handler.service.GetCampaign(c.Request.Context(), c.Param("workspace_id"), c.Param("campaign_id"))
	common.Respond(c, item, err, http.StatusOK)
}

func (handler *Handler) UpdateCampaign(c *gin.Context) {
	var input UpdateCampaignInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.UpdateCampaign(c.Request.Context(), c.Param("workspace_id"), c.Param("campaign_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusOK)
}

func (handler *Handler) CreateChannelPackage(c *gin.Context) {
	var input CreateChannelPackageInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateChannelPackage(c.Request.Context(), c.Param("workspace_id"), c.Param("campaign_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) ListPublishingTasks(c *gin.Context) {
	items, err := handler.service.ListPublishingTasks(c.Request.Context(), c.Param("workspace_id"), c.Param("campaign_id"))
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) CreatePublishingTask(c *gin.Context) {
	var input CreatePublishingTaskInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreatePublishingTask(c.Request.Context(), c.Param("workspace_id"), c.Param("campaign_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) UpdatePublishingTask(c *gin.Context) {
	var input UpdatePublishingTaskInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.UpdatePublishingTask(c.Request.Context(), c.Param("workspace_id"), c.Param("campaign_id"), c.Param("task_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusOK)
}
