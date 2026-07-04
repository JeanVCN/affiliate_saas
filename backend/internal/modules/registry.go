package modules

import (
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/affiliate"
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/analytics"
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/identity"
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/linktracking"
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/marketplace"
	"github.com/JeanVCN/affiliate_saas/backend/internal/modules/product"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	DB           *pgxpool.Pool
	Identity     *identity.Service
	Marketplace  *marketplace.Service
	Product      *product.Service
	Affiliate    *affiliate.Service
	LinkTracking *linktracking.Service
	Analytics    *analytics.Service
}

func (deps Dependencies) HasServices() bool {
	return deps.DB != nil ||
		deps.Identity != nil ||
		deps.Marketplace != nil ||
		deps.Product != nil ||
		deps.Affiliate != nil ||
		deps.LinkTracking != nil ||
		deps.Analytics != nil
}

func RegisterRoutes(router *gin.Engine, deps Dependencies) {
	deps = withPostgresDefaults(deps)

	api := router.Group("/api/v1")
	if deps.Identity != nil {
		identity.RegisterRoutes(api, deps.Identity)
	}
	if deps.Marketplace != nil {
		marketplace.RegisterRoutes(api, deps.Marketplace)
	}
	if deps.Product != nil {
		product.RegisterRoutes(api, deps.Product)
	}
	if deps.Affiliate != nil {
		affiliate.RegisterRoutes(api, deps.Affiliate)
	}
	if deps.LinkTracking != nil {
		linktracking.RegisterRoutes(api, deps.LinkTracking)
		linktracking.RegisterPublicRoutes(router, deps.LinkTracking)
	}
	if deps.Analytics != nil {
		analytics.RegisterRoutes(api, deps.Analytics)
	}
}

func withPostgresDefaults(deps Dependencies) Dependencies {
	if deps.DB == nil {
		return deps
	}
	if deps.Identity == nil {
		deps.Identity = identity.NewService(identity.NewPostgresRepository(deps.DB))
	}
	if deps.Marketplace == nil {
		deps.Marketplace = marketplace.NewService(marketplace.NewPostgresRepository(deps.DB))
	}
	if deps.Product == nil {
		deps.Product = product.NewService(product.NewPostgresRepository(deps.DB))
	}
	if deps.Affiliate == nil {
		deps.Affiliate = affiliate.NewService(affiliate.NewPostgresRepository(deps.DB))
	}
	if deps.LinkTracking == nil {
		deps.LinkTracking = linktracking.NewService(linktracking.NewPostgresRepository(deps.DB))
	}
	if deps.Analytics == nil {
		deps.Analytics = analytics.NewService(analytics.NewPostgresRepository(deps.DB))
	}
	return deps
}
