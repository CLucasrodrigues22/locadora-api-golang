package brands

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/exceptions"
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/repositories/brands"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHandler(ctx *gin.Context) {
	req := schemas.BrandPayload{}

	if exceptions.ValidatePayload(ctx, ctx.ShouldBind(&req), http.StatusBadRequest, "Invalid payload") {
		return
	}

	brand, err := brands.CreateBrand(req)

	if err != nil {
		handlers.SendError(ctx, http.StatusInternalServerError, err.Error())
	}

	handlers.SendSuccess(ctx, "Brand created", brand, http.StatusCreated)
}
