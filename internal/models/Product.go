package models

import "time"

type Product struct {
	ID          int         `json:"id"`
	SellerID    int         `json:"seller_id"`
	Name        string      `json:"name" validate:"required,min=3"`
	Description Description `json:"description"`
	Price       float64     `json:"price"`
	Quantity    int         `json:"quantity"`
	Brand       string      `json:"brand"`
	Sale        float64     `json:"sale"`
	Image       string      `json:"image"`
	CategoryID  int         `json:"category_id"`
	Article     string      `json:"article"`
	Rating      float64     `json:"rating"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type Description struct {
	Color    string `json:"color"`
	Size     string `json:"size"`
	Material string `json:"material"`
	Describe string `json:"describe"`
}
