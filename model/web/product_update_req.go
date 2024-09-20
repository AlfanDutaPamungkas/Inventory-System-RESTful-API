package web

import "time"

type ProductUpdateReq struct {
	SKU         string    `schema:"sku"`
	Name        string    `schema:"product_name"`
	Brand       string    `schema:"product_brand"`
	Category    string    `schema:"category"`
	Price       int       `schema:"price"`
	Amount      int       `schema:"stock_amount"`
	ExpiredDate *time.Time `schema:"expired_at"`
}
