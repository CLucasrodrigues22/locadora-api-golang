package main

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/configs"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/CLucasrodrigues22/api-locadora/internal/routes"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
)

var (
	logger *logs.Logger
)

func main() {
	logger = utils.GetLogger("main")

	err := configs.Init()

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
	}

	routes.InitializeRouter(logger)
}
