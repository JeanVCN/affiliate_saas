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
	api.GET("/workspaces/:workspace_id/analytics/overview", handler.Overview)
	api.GET("/workspaces/:workspace_id/analytics/clicks", handler.ClickMetrics)
	api.GET("/workspaces/:workspace_id/analytics/top-products", handler.TopProducts)
	api.POST("/workspaces/:workspace_id/conversion-imports", handler.CreateConversionImport)
	api.POST("/workspaces/:workspace_id/conversion-imports/:import_id/rows", handler.CreateConversionImportRow)
	api.GET("/workspaces/:workspace_id/conversion-imports/:import_id", handler.GetConversionImport)
}

func RegisterWorkspaceRoutes(workspace gin.IRouter, service *Service) {
	handler := NewHandler(service)
	workspace.GET("/analytics/overview", handler.Overview)
	workspace.GET("/analytics/clicks", handler.ClickMetrics)
	workspace.GET("/analytics/top-products", handler.TopProducts)
	workspace.POST("/conversion-imports", handler.CreateConversionImport)
	workspace.POST("/conversion-imports/:import_id/rows", handler.CreateConversionImportRow)
	workspace.GET("/conversion-imports/:import_id", handler.GetConversionImport)
}

func (handler *Handler) Overview(c *gin.Context) {
	item, err := handler.service.Overview(c.Request.Context(), c.Param("workspace_id"))
	common.Respond(c, item, err, http.StatusOK)
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

func (handler *Handler) TopProducts(c *gin.Context) {
	items, err := handler.service.TopProducts(c.Request.Context(), c.Param("workspace_id"))
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) CreateConversionImport(c *gin.Context) {
	var input CreateConversionImportInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateConversionImport(c.Request.Context(), c.Param("workspace_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) CreateConversionImportRow(c *gin.Context) {
	var input CreateConversionImportRowInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateConversionImportRow(c.Request.Context(), c.Param("workspace_id"), c.Param("import_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) GetConversionImport(c *gin.Context) {
	item, err := handler.service.GetConversionImport(c.Request.Context(), c.Param("workspace_id"), c.Param("import_id"))
	common.Respond(c, item, err, http.StatusOK)
}
