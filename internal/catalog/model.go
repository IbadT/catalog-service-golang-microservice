package catalog

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title        string    `json:"title" gorm:"type:varchar(255);not null"`
	Price        float32   `json:"price"`
	CurrentPrice float32   `json:"current_price"`
	OldPrice     float32   `json:"old_price"`
	Descripton   string    `json:"description" gorm:"type:varchar(255);not null"`
	Rating       float32   `json:"rating"`
	ImageUrl     string    `json:"image_url" gorm:"type:varchar(255);not null"`
	Mimetype     string    `json:"mimetype" gorm:"type:varchar(255);not null"`
	Count        uint      `json:"count"`
	Code         string    `json:"code"`
	CategoryID   []Category
}

type Category struct {
	gorm.Model

	ID           uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CategoryName string    `json:"category_name" gorm:"type:varchar(255);not null"`
}
