package handlers

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/gin-gonic/gin"
)

func DashboardHandler(ctx *gin.Context) {
	var dashboard []schemas.DashboardResponse

	sendSuccess(ctx, "Dashboard", dashboard)
}
