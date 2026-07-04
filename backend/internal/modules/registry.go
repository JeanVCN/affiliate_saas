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

func RegisterRoutes(router *gin.Engine, deps Dependencies, appEnv string) {
	deps = withPostgresDefaults(deps)

	api := router.Group("/api/v1")
	if deps.Identity != nil {
		protected := api.Group("")
		protected.Use(identity.AuthMiddleware(deps.Identity))
		workspace := api.Group("/workspaces/:workspace_id")
		workspace.Use(identity.AuthMiddleware(deps.Identity), identity.WorkspaceRoleMiddleware(deps.Identity, identity.RoleMember))

		identity.RegisterAuthRoutes(api, protected, deps.Identity, appEnv)
		identity.RegisterRoutes(protected, deps.Identity)

		if deps.Marketplace != nil {
			marketplace.RegisterWorkspaceRoutes(workspace, deps.Marketplace)
		}
		if deps.Product != nil {
			product.RegisterWorkspaceRoutes(workspace, deps.Product)
		}
		if deps.Affiliate != nil {
			affiliate.RegisterWorkspaceRoutes(workspace, deps.Affiliate)
		}
		if deps.LinkTracking != nil {
			linktracking.RegisterWorkspaceRoutes(workspace, deps.LinkTracking)
		}
		if deps.Analytics != nil {
			analytics.RegisterWorkspaceRoutes(workspace, deps.Analytics)
		}
	}
	if deps.Marketplace != nil {
		protected := api.Group("")
		if deps.Identity != nil {
			protected.Use(identity.AuthMiddleware(deps.Identity))
		}
		marketplace.RegisterRoutes(protected, deps.Marketplace)
	}
	if deps.Identity == nil {
		if deps.Product != nil {
			product.RegisterRoutes(api, deps.Product)
		}
		if deps.Affiliate != nil {
			affiliate.RegisterRoutes(api, deps.Affiliate)
		}
		if deps.LinkTracking != nil {
			linktracking.RegisterRoutes(api, deps.LinkTracking)
		}
		if deps.Analytics != nil {
			analytics.RegisterRoutes(api, deps.Analytics)
		}
	}
	if deps.LinkTracking != nil {
		linktracking.RegisterPublicRoutes(router, deps.LinkTracking)
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
