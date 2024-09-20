package web

import "time"

type ProductResponse struct {
	SKU         string    `json:"sku"`
	Name        string    `json:"product_name"`
	Brand       string    `json:"product_brand"`
	Category    string    `json:"category"`
	Price       int       `json:"price"`
	ImageUrl    string    `json:"image_url"`
	Amount      int       `json:"stock_amount"`
	ExpiredDate *time.Time `json:"expired_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
}
