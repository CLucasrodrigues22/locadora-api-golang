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
	Name  string                `form:"name" binding:"required"`
	Icon  *multipart.FileHeader `form:"icon" binding:"omitempty"`
	Image string                `form:"-"`
}

type ListBrandsResponse struct {
	Data    []BrandResponse `json:"data"`
	Status  int             `json:"status"`
	Message string          `json:"message"`
}

func (b *Brand) GetImagePath() string {
	return b.Icon
}

func (b *Brand) SetImagePath(path string) {
	b.Icon = path
}
