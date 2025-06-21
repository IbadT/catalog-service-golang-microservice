package domain

import "github.com/google/uuid"

type Category struct {
	ID           uuid.UUID `json:"id" validate:"uuid"`
	CategoryName string    `json:"category_name"`
}

func NewCategory(categoryName string) *Category {
	uid := uuid.New()
	return &Category{
		ID:           uid,
		CategoryName: categoryName,
	}
}
