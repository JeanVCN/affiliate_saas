package analytics

import (
	"errors"
	"io"
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
	api.POST("/workspaces/:workspace_id/conversion-imports/:import_id/csv-rows", handler.CreateConversionImportCSVRows)
	api.POST("/workspaces/:workspace_id/conversion-imports/:import_id/csv-upload", handler.UploadConversionImportCSV)
	api.GET("/workspaces/:workspace_id/conversion-imports/:import_id/reconciliation", handler.ReconciliationSummary)
	api.PATCH("/workspaces/:workspace_id/conversion-imports/:import_id/rows/:row_id", handler.UpdateConversionImportRow)
	api.GET("/workspaces/:workspace_id/conversion-imports/:import_id", handler.GetConversionImport)
}

func RegisterWorkspaceRoutes(workspace gin.IRouter, service *Service) {
	handler := NewHandler(service)
	workspace.GET("/analytics/overview", handler.Overview)
	workspace.GET("/analytics/clicks", handler.ClickMetrics)
	workspace.GET("/analytics/top-products", handler.TopProducts)
	workspace.POST("/conversion-imports", handler.CreateConversionImport)
	workspace.POST("/conversion-imports/:import_id/rows", handler.CreateConversionImportRow)
	workspace.POST("/conversion-imports/:import_id/csv-rows", handler.CreateConversionImportCSVRows)
	workspace.POST("/conversion-imports/:import_id/csv-upload", handler.UploadConversionImportCSV)
	workspace.GET("/conversion-imports/:import_id/reconciliation", handler.ReconciliationSummary)
	workspace.PATCH("/conversion-imports/:import_id/rows/:row_id", handler.UpdateConversionImportRow)
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

func (handler *Handler) CreateConversionImportCSVRows(c *gin.Context) {
	var input CreateConversionImportCSVInput
	if !common.BindJSON(c, &input) {
		return
	}
	items, err := handler.service.CreateConversionImportCSVRows(c.Request.Context(), c.Param("workspace_id"), c.Param("import_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, gin.H{"rows": items}, err, http.StatusCreated)
}

func (handler *Handler) UploadConversionImportCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.RespondValidationError(c, errors.New("file is required"))
		return
	}
	if file.Size > 2*1024*1024 {
		common.RespondValidationError(c, errors.New("file must be 2MB or smaller"))
		return
	}
	opened, err := file.Open()
	if err != nil {
		common.Respond(c, nil, err, http.StatusCreated)
		return
	}
	defer opened.Close()

	items, err := handler.service.CreateConversionImportCSVRowsFromReader(c.Request.Context(), c.Param("workspace_id"), c.Param("import_id"), io.LimitReader(opened, 2*1024*1024+1))
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, gin.H{"rows": items}, err, http.StatusCreated)
}

func (handler *Handler) GetConversionImport(c *gin.Context) {
	item, err := handler.service.GetConversionImport(c.Request.Context(), c.Param("workspace_id"), c.Param("import_id"))
	common.Respond(c, item, err, http.StatusOK)
}

func (handler *Handler) ReconciliationSummary(c *gin.Context) {
	item, err := handler.service.ReconciliationSummary(c.Request.Context(), c.Param("workspace_id"), c.Param("import_id"))
	common.Respond(c, item, err, http.StatusOK)
}

func (handler *Handler) UpdateConversionImportRow(c *gin.Context) {
	var input UpdateConversionImportRowInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.UpdateConversionImportRow(c.Request.Context(), c.Param("workspace_id"), c.Param("import_id"), c.Param("row_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusOK)
}
