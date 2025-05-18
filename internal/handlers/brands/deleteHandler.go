package brands

import (
	"errors"
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/repositories/brands"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		handlers.SendError(ctx, http.StatusBadRequest, "Invalid ID")
		return
	}
	var brand schemas.Brand
	if err := handlers.Db.First(&brand, id).Error; err != nil {
		handlers.SendError(ctx, http.StatusNotFound, "Brand not found")
		return
	}

	err := handlers.DeleteFile(ctx, &brand)
	if err != nil {
		handlers.SendError(ctx, http.StatusInternalServerError, "Error deleting file")
		return
	}

	err = brands.DeleteBrand(id)

	if err != nil {
		if errors.Is(err, brands.ErrBrandNotFound) {
			handlers.SendError(ctx, http.StatusNotFound, "Brand not found")
			return
		}
		handlers.SendError(ctx, http.StatusInternalServerError, "Error deleting brand")
		return
	}

	handlers.SendSuccess(ctx, "Brand deleted", nil, http.StatusOK)
}
