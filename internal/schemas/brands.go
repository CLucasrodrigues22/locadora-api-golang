package schemas

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name string
	Icon string
}

type BrandResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deletedAt,omitzero"`
}

type BrandPayload struct {
	Name string                `form:"name" binding:"required"`
	Icon *multipart.FileHeader `form:"icon" binding:"required"`
}

type ListBrandsResponse struct {
	Message string          `json:"message"`
	Data    []BrandResponse `json:"data"`
}
