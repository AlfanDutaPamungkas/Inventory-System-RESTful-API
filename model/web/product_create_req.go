package web

import "time"

type ProductCreateReq struct {
	SKU         string     `schema:"sku" validate:"required,max=100"`
	Name        string     `schema:"product_name" validate:"required,max=200"`
	Brand       string     `schema:"product_brand" validate:"required,max=200"`
	Category    string     `schema:"category" validate:"required,max=100"`
	Price       int        `schema:"price" validate:"required,numeric"`
	Amount      int        `schema:"stock_amount" validate:"required,numeric"`
	ExpiredDate *time.Time `schema:"expired_at"`
}
