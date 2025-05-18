package routes

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/gin-gonic/gin"
)

var logger *logs.Logger

func InitializeRouter(l *logs.Logger, router *gin.Engine) {
	logger = l

	initializeRoutes(router)
}
