package routes

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	handlers.InitHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", handlers.DashboardHandler)
	}

	logger.Infof("Server started on port %s", utils.GetEnv("PORT"))

	if err := router.Run(":" + utils.GetEnv("PORT")); err != nil {
		logger.Error("Failed to start the server: " + err.Error())
	}
}
