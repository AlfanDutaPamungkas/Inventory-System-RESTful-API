package web

type StockAmountReq struct {
	SKU    string `schema:"sku"`
	Amount int    `schema:"stock_amount" validate:"required,numeric"`
}
