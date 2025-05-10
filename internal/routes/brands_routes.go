package routes

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers/brands"
	"github.com/gin-gonic/gin"
)

func BrandsRoutes(router *gin.RouterGroup) {
	brandsGroup := router.Group("/brands")
	{
		brandsGroup.GET("/", brands.IndexHandler)
		brandsGroup.GET("/:id", brands.ShowHandler)
		brandsGroup.POST("/", brands.CreateHandler)
		brandsGroup.PUT("/:id", brands.UpdateHandler)
		brandsGroup.DELETE("/:id", brands.DeleteHandler)
	}
}
