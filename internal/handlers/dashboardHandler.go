package handlers

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DashboardHandler(ctx *gin.Context) {
	var dashboard []schemas.DashboardResponse

	SendSuccess(ctx, "Dashboard", dashboard, http.StatusOK)
}
