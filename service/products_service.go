package service

import (
	"context"
	"inventory-system-api/model/web"
	"mime/multipart"
)

type ProductsService interface {
	CreateProductService(ctx context.Context, request web.ProductCreateReq, file multipart.File, fileHeader *multipart.FileHeader) web.ProductResponse
	FindAllService(ctx context.Context) []web.ProductResponse
	FindBySKUService(ctx context.Context, SKU string) web.ProductResponse
	UpdateProductService(ctx context.Context, request web.ProductUpdateReq) web.ProductResponse
	StockOutService(ctx context.Context, request web.StockAmountReq) web.ProductResponse
	StockInService(ctx context.Context, request web.StockAmountReq) web.ProductResponse
	UpdateImgUrlService(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, SKU string) web.ProductResponse
	NullifyExpiredDateService(ctx context.Context, SKU string) web.ProductResponse
	Delete(ctx context.Context, SKU string) string
}
