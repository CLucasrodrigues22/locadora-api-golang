package brands

import (
	"errors"
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"gorm.io/gorm"
	"time"
)

func UpdateBrand(id string, req *schemas.BrandPayload) (*schemas.BrandResponse, error) {
	brand := &schemas.Brand{}

	if err := handlers.Db.First(brand, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBrandNotFound
		}
		return nil, err
	}

	brand.Name = req.Name

	if err := handlers.Db.Save(brand).Error; err != nil {
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
