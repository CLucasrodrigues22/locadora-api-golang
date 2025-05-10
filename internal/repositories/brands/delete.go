package brands

import (
	"errors"
	"github.com/CLucasrodrigues22/api-locadora/internal/handlers"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"gorm.io/gorm"
)

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
