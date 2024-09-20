package web

import "time"

type ProductCreateReq struct {
	SKU         string    `schema:"sku,required"`
	Name        string    `schema:"product_name,required"`
	Brand       string    `schema:"product_brand,required"`
	Category    string    `schema:"category,required"`
	Price       int       `schema:"price,required"`
	Amount      int       `schema:"stock_amount,required"`
	ExpiredDate *time.Time `schema:"expired_at"`
}
