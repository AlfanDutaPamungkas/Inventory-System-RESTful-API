package web

import "time"

type ProductUpdateReq struct {
	SKU         string     `schema:"sku"`
	Name        string     `schema:"product_name" validate:"max=200"`
	Brand       string     `schema:"product_brand" validate:"max=200"`
	Category    string     `schema:"category" validate:"max=100"`
	Price       int        `schema:"price" validate:"numeric"`
	Amount      int        `schema:"stock_amount" validate:"numeric"`
	ExpiredDate *time.Time `schema:"expired_at"`
}
