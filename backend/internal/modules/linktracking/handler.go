package linktracking

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net"
	"net/http"
	"strings"

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
	api.POST("/workspaces/:workspace_id/links/:link_id/short-links", handler.CreateShortLink)
}

func RegisterWorkspaceRoutes(workspace gin.IRouter, service *Service) {
	handler := NewHandler(service)
	workspace.POST("/links/:link_id/short-links", handler.CreateShortLink)
}

func RegisterPublicRoutes(router gin.IRouter, service *Service) {
	handler := NewHandler(service)
	router.GET("/r/:slug", handler.Redirect)
}

func (handler *Handler) CreateShortLink(c *gin.Context) {
	var input CreateShortLinkInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateShortLink(c.Request.Context(), c.Param("workspace_id"), c.Param("link_id"), input)
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) Redirect(c *gin.Context) {
	resolved, err := handler.service.ResolveShortLink(c.Request.Context(), c.Param("slug"))
	if err != nil {
		common.Respond(c, nil, err, http.StatusOK)
		return
	}

	click := RecordClickInput{
		WorkspaceID:     resolved.WorkspaceID,
		ShortLinkID:     resolved.ShortLinkID,
		AffiliateLinkID: resolved.AffiliateLinkID,
		ProductID:       resolved.ProductID,
		Referrer:        c.GetHeader("Referer"),
		UserAgent:       c.GetHeader("User-Agent"),
		IPHash:          hashIP(c.ClientIP()),
		UTMSource:       resolved.UTMSource,
		UTMMedium:       resolved.UTMMedium,
		UTMCampaign:     resolved.UTMCampaign,
	}
	if err := handler.service.RecordClick(c.Request.Context(), click); err != nil {
		log.Printf("record click failed slug=%s: %v", c.Param("slug"), err)
	}
	c.Redirect(http.StatusFound, resolved.DestinationURL)
}

func hashIP(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if parsed := net.ParseIP(value); parsed != nil {
		value = parsed.String()
	}
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:])
}
