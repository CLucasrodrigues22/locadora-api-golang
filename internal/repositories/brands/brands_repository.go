package brands

import (
	"errors"
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"gorm.io/gorm"
	"time"
)

var (
	ErrBrandNotFound = errors.New("brand not found")
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

func CreateBrand(req schemas.BrandPayload) (*schemas.BrandResponse, error) {
	brand := schemas.Brand{
		Name: req.Name,
		Icon: req.Image,
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

func ShowBrand(id string) (*schemas.BrandResponse, error) {
	brand := &schemas.Brand{}

	if err := handlers.Db.First(brand, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBrandNotFound
		}
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

func UpdateBrand(id string, req *schemas.BrandPayload) (*schemas.BrandResponse, error) {
	brand := &schemas.Brand{}

	if err := handlers.Db.First(brand, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBrandNotFound
		}
		return nil, err
	}

	brand.Name = req.Name
	brand.Icon = req.Image

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

func DeleteBrand(id string) error {
	var brand schemas.Brand

	if err := handlers.Db.First(&brand, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrBrandNotFound
		}
		return err
	}

	return handlers.Db.Delete(&brand).Error
}
