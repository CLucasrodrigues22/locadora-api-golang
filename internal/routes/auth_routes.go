package routes

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	loginGroup := router.Group("/auth/sso")
	{
		// google sso
		loginGroup.GET("/google/login", auth.GoogleLoginHandler)
		loginGroup.GET("/google/callback", auth.GoogleCallbackHandler)

		// local login
	}
}
