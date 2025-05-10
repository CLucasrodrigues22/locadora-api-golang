package brands

import (
	"errors"
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/repositories/brands"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		handlers.SendError(ctx, http.StatusBadRequest, "Brand not found")
		return
	}

	brand, err := brands.ShowBrand(id)

	if err != nil {
		if errors.Is(err, brands.ErrBrandNotFound) {
			handlers.SendError(ctx, http.StatusNotFound, "Brand not found")
			return
		}
		handlers.SendError(ctx, http.StatusInternalServerError, "Error showing brand")
		return
	}

	handlers.SendSuccess(ctx, "Brand found", brand, http.StatusOK)
}
