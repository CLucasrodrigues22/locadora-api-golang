package routes

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	loginGroup := router.Group("/auth/google")
	{
		loginGroup.GET("/login", auth.GoogleLoginHandler)
		loginGroup.GET("/callback", auth.GoogleCallbackHandler)
	}
}
