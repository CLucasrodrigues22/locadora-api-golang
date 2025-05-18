package routes

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	handlers.InitHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", handlers.DashboardHandler)
		AuthRoutes(v1)
		BrandsRoutes(v1)
	}
}
