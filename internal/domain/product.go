package domain

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id" validate:"uuid"`
	Title       string    `json:"title"`
	Price       float32   `json:"price"`
	OldPrice    float32   `json:"old_price"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	ImageUrl    string    `json:"image_url"`
	Mimetype    string    `json:"mimetype"`
	Count       uint      `json:"count"`
	Code        string    `json:"code"`
	CategoryId  string    `json:"categor_id" validate:"uuid"`
}

func NewProduct(title, code, description, categoryId string, price float32, count int32) *Product {
	uid := uuid.New()

	// добавить validate
	return &Product{
		ID:          uid,
		Title:       title,
		Description: description,
		Code:        code,
		Count:       uint(count),
		Price:       price,
		OldPrice:    price,
		CategoryId:  categoryId,
	}
}

func NewProductUUID(uid uuid.UUID) *Product {
	return &Product{
		ID: uid,
	}
}

// func ChangeProductTitle(title string) *Product {
// 	return &Product{
// 		Title: title,
// 	}
// }

// func ChangeProductPrice(price, oldPrice float32) *Product {
// 	return &Product{
// 		Price:    price,
// 		OldPrice: oldPrice,
// 	}
// }
