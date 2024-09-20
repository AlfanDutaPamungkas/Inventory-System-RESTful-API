package helper

import (
	"inventory-system-api/model/domain"
	"inventory-system-api/model/web"
	"time"
)

func ToProductResponse(product domain.Products, stock domain.ProductStock) web.ProductResponse {
	now := time.Now()

	productDiff := now.Sub(product.UpdatedAt)
	stockDiff := now.Sub(stock.UpdatedAt)

	var updatedAt time.Time
	if productDiff < stockDiff {
		updatedAt = product.UpdatedAt
	} else {
		updatedAt = stock.UpdatedAt
	}

	return web.ProductResponse{
		SKU:         product.SKU,
		Name:        product.Name,
		Brand:       product.Brand,
		Category:    product.Category,
		Price:       product.Price,
		ImageUrl:    product.ImageUrl,
		Amount:      stock.Amount,
		ExpiredDate: stock.ExpiredDate,
		CreatedAt:   product.CreatedAt,
		UpdateAt:    updatedAt,
	}
}
