package domain

import "github.com/google/uuid"

type Category struct {
	ID           uuid.UUID `json:"id"`
	CategoryName string    `json:"category_name"`
}
