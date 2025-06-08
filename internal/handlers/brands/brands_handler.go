package brands

import (
	"errors"
	"github.com/CLucasrodrigues22/api-locadora/internal/exceptions"
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	brandsRepository "github.com/CLucasrodrigues22/api-locadora/internal/repositories/brands"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
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

func CreateHandler(ctx *gin.Context) {
	req := schemas.BrandPayload{}

	if exceptions.ValidatePayload(ctx, ctx.ShouldBind(&req), http.StatusBadRequest, "Invalid payload") {
		return
	}

	file, err := handlers.SaveFile(ctx, req.Icon)

	if err != nil {
		handlers.SendError(ctx, http.StatusInternalServerError, "Error during file upload")
		return
	}

	req.Image = file

	brand, err := brandsRepository.CreateBrand(req)

	if err != nil {
		handlers.SendError(ctx, http.StatusInternalServerError, err.Error())
	}

	handlers.SendSuccess(ctx, "Brand created", brand, http.StatusCreated)
}

func ShowHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		handlers.SendError(ctx, http.StatusBadRequest, "Brand not found")
		return
	}

	brand, err := brandsRepository.ShowBrand(id)

	if err != nil {
		if errors.Is(err, brandsRepository.ErrBrandNotFound) {
			handlers.SendError(ctx, http.StatusNotFound, "Brand not found")
			return
		}
		handlers.SendError(ctx, http.StatusInternalServerError, "Error showing brand")
		return
	}

	handlers.SendSuccess(ctx, "Brand found", brand, http.StatusOK)
}

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

	schema, err := brandsRepository.UpdateBrand(id, &req)
	if err != nil {
		if errors.Is(err, brandsRepository.ErrBrandNotFound) {
			handlers.SendError(ctx, http.StatusNotFound, "Brand not found")
			return
		}
		handlers.SendError(ctx, http.StatusInternalServerError, "Error updating brand")
		return
	}

	handlers.SendSuccess(ctx, "Brand updated", schema, http.StatusOK)
}

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

	err = brandsRepository.DeleteBrand(id)

	if err != nil {
		if errors.Is(err, brandsRepository.ErrBrandNotFound) {
			handlers.SendError(ctx, http.StatusNotFound, "Brand not found")
			return
		}
		handlers.SendError(ctx, http.StatusInternalServerError, "Error deleting brand")
		return
	}

	handlers.SendSuccess(ctx, "Brand deleted", nil, http.StatusOK)
}
