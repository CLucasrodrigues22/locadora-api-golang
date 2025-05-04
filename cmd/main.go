package main

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/configs"
	"github.com/CLucasrodrigues22/api-locadora/internal/configs/logs"
)

var (
	logger *logs.Logger
)

func main() {
	logger = configs.GetLogger("main")
}
