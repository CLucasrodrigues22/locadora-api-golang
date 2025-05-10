package brands

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"time"
)

func CreateBrand(req schemas.BrandPayload) (*schemas.BrandResponse, error) {
	brand := schemas.Brand{
		Name: req.Name,
		Icon: "Icon/brand",
	}

	if err := handlers.Db.Create(&brand).Error; err != nil {
		return nil, err
	}

	res := schemas.BrandResponse{
		ID:        brand.ID,
		Name:      brand.Name,
		Icon:      brand.Icon,
		CreatedAt: brand.CreatedAt,
		UpdatedAt: brand.UpdatedAt,
		DeletedAt: time.Time{},
	}

	return &res, nil
}
