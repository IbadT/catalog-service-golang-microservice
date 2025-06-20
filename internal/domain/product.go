package domain

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Price       int32     `json:"price"`
	OldPrice    int32     `json:"old_price"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	ImageUrl    string    `json:"image_url"`
	Mimetype    string    `json:"mimetype"`
	Count       int32     `json:"count"`
	Code        string    `json:"code"`
	CategoryId  string    `json:"categor_id"`
}
