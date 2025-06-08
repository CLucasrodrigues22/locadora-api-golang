package main

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/bootstrap"
	"github.com/CLucasrodrigues22/api-locadora/internal/configs"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/CLucasrodrigues22/api-locadora/internal/routes"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
	"github.com/gin-gonic/gin"
)

var (
	logger *logs.Logger
)

func main() {
	logger = bootstrap.GetLogger("main")

	err := configs.Init()

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
	}

	router := gin.Default()
	routes.InitializeRouter(logger, router)

	logger.Infof("Server started on port %s", utils.GetEnv("PORT", logger))

	if err := router.Run(":" + utils.GetEnv("PORT", logger)); err != nil {
		logger.Error("Failed to start the server: " + err.Error())
	}
}
