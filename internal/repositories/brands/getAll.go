package brands

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"time"
)

func GetAllBrands() ([]schemas.BrandResponse, error) {
	var brands []schemas.Brand

	if err := handlers.Db.Find(&brands).Error; err != nil {
		return nil, err
	}

	var response []schemas.BrandResponse

	for _, brand := range brands {
		brandResponse := schemas.BrandResponse{
			ID:        brand.ID,
			Name:      brand.Name,
			Icon:      brand.Icon,
			CreatedAt: brand.CreatedAt,
			UpdatedAt: brand.UpdatedAt,
			DeletedAt: time.Time{},
		}
		response = append(response, brandResponse)
	}

	return response, nil
}
