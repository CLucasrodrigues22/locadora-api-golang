package brands

import (
	"errors"
	"github.com/CLucasrodrigues22/api-locadora/internal/exceptions"
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/repositories/brands"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

var brand schemas.Brand

func UpdateHandler(ctx *gin.Context) {
	req := schemas.BrandPayload{}

	if exceptions.ValidatePayload(ctx, ctx.ShouldBind(&req), 400, "Invalid payload") {
		return
	}

	id := ctx.Param("id")
	if id == "" {
		handlers.SendError(ctx, http.StatusBadRequest, "Brand not found")
		return
	}

	var brand schemas.Brand
	if err := handlers.Db.First(&brand, id).Error; err != nil {
		handlers.SendError(ctx, http.StatusNotFound, "Brand not found")
		return
	}

	file, err := handlers.UpdateFile(ctx, req.Icon, &brand)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, "Error uploading file")
		return
	}

	req.Image = file

	schema, err := brands.UpdateBrand(id, &req)
	if err != nil {
		if errors.Is(err, brands.ErrBrandNotFound) {
			handlers.SendError(ctx, http.StatusNotFound, "Brand not found")
			return
		}
		handlers.SendError(ctx, http.StatusInternalServerError, "Error updating brand")
		return
	}

	handlers.SendSuccess(ctx, "Brand updated", schema, http.StatusOK)
}
