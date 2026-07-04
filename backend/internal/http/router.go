package httpapi

import (
	"net/http"

	"github.com/JeanVCN/affiliate_saas/backend/internal/modules"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	AppEnv  string
	DB      *pgxpool.Pool
	Modules modules.Dependencies
}

func NewRouter(deps Dependencies) *gin.Engine {
	if deps.AppEnv == "test" {
		gin.SetMode(gin.TestMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"db": gin.H{
				"configured": deps.DB != nil,
			},
		})
	})

	moduleDeps := deps.Modules
	if moduleDeps.DB == nil {
		moduleDeps.DB = deps.DB
	}
	if moduleDeps.HasServices() {
		modules.RegisterRoutes(router, moduleDeps)
	}

	return router
}
