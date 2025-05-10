package brands

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	brandsRepository "github.com/CLucasrodrigues22/api-locadora/internal/repositories/brands"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(ctx *gin.Context) {
	brands, err := brandsRepository.GetAllBrands()

	if err != nil {
		handlers.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handlers.SendSuccess(ctx, "Brands fetched", brands, http.StatusOK)
}
