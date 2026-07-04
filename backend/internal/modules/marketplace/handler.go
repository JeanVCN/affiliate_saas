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
}

func RegisterWorkspaceRoutes(workspace gin.IRouter, service *Service) {
	handler := NewHandler(service)
	workspace.GET("/programs", handler.ListWorkspacePrograms)
	workspace.POST("/programs", handler.EnableWorkspaceProgram)
	workspace.GET("/programs/:program_id/policy-notes", handler.ListProgramPolicyNotes)
	workspace.POST("/programs/:program_id/policy-notes", handler.CreateProgramPolicyNote)
	workspace.PATCH("/programs/:program_id/policy-notes/:note_id", handler.UpdateProgramPolicyNote)
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

func (handler *Handler) CreateProgramPolicyNote(c *gin.Context) {
	var input CreateProgramPolicyNoteInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateProgramPolicyNote(c.Request.Context(), c.Param("workspace_id"), c.Param("program_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) ListProgramPolicyNotes(c *gin.Context) {
	items, err := handler.service.ListProgramPolicyNotes(c.Request.Context(), c.Param("workspace_id"), c.Param("program_id"))
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) UpdateProgramPolicyNote(c *gin.Context) {
	var input UpdateProgramPolicyNoteInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.UpdateProgramPolicyNote(c.Request.Context(), c.Param("workspace_id"), c.Param("program_id"), c.Param("note_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusOK)
}
