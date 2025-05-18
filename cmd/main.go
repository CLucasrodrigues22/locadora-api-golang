package main

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/common"
	"github.com/CLucasrodrigues22/api-locadora/internal/configs"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/CLucasrodrigues22/api-locadora/internal/routes"
	"github.com/gin-gonic/gin"
)

var (
	logger *logs.Logger
)

func main() {
	logger = common.GetLogger("main")

	err := configs.Init()

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
	}

	router := gin.Default()
	routes.InitializeRouter(logger, router)

	logger.Infof("Server started on port %s", common.GetEnv("PORT"))

	if err := router.Run(":" + common.GetEnv("PORT")); err != nil {
		logger.Error("Failed to start the server: " + err.Error())
	}
}
