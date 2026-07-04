package product

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
	api.GET("/workspaces/:workspace_id/products", handler.ListProducts)
	api.POST("/workspaces/:workspace_id/products", handler.CreateProduct)
	api.GET("/workspaces/:workspace_id/products/:product_id", handler.GetProduct)
	api.POST("/workspaces/:workspace_id/products/:product_id/offers", handler.CreateOffer)
}

func RegisterWorkspaceRoutes(workspace gin.IRouter, service *Service) {
	handler := NewHandler(service)
	workspace.GET("/products", handler.ListProducts)
	workspace.POST("/products", handler.CreateProduct)
	workspace.GET("/products/:product_id", handler.GetProduct)
	workspace.POST("/products/:product_id/offers", handler.CreateOffer)
}

func (handler *Handler) ListProducts(c *gin.Context) {
	items, err := handler.service.ListProducts(c.Request.Context(), c.Param("workspace_id"))
	common.Respond(c, items, err, http.StatusOK)
}

func (handler *Handler) CreateProduct(c *gin.Context) {
	var input CreateProductInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateProduct(c.Request.Context(), c.Param("workspace_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}

func (handler *Handler) GetProduct(c *gin.Context) {
	item, err := handler.service.GetProduct(c.Request.Context(), c.Param("workspace_id"), c.Param("product_id"))
	common.Respond(c, item, err, http.StatusOK)
}

func (handler *Handler) CreateOffer(c *gin.Context) {
	var input CreateOfferInput
	if !common.BindJSON(c, &input) {
		return
	}
	item, err := handler.service.CreateOffer(c.Request.Context(), c.Param("workspace_id"), c.Param("product_id"), input)
	if common.IsValidationError(err) {
		common.RespondValidationError(c, err)
		return
	}
	common.Respond(c, item, err, http.StatusCreated)
}
