package helper

import (
	"inventory-system-api/model/domain"
	"inventory-system-api/model/web"
)

func ToProductResponse(product domain.Products, stock domain.ProductStock) web.ProductResponse {
	return web.ProductResponse{
		SKU:         product.SKU,
		Name:        product.Name,
		Brand:       product.Brand,
		Category:    product.Category,
		Price:       product.Price,
		ImageUrl:    product.ImageUrl,
		Amount:      product.Amount,
		ExpiredDate: product.ExpiredDate,
		CreatedAt:   product.CreatedAt,
		UpdateAt:    product.UpdatedAt,
	}
}
