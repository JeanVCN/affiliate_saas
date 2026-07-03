package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	AppEnv string
	DB     *pgxpool.Pool
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

	return router
}
