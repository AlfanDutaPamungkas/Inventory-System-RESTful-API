package repository

import (
	"context"
	"database/sql"
	"inventory-system-api/model/domain"
)

type ProductsRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products 
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Products
	FindBySKU(ctx context.Context, tx *sql.Tx, SKU string) (domain.Products, error)
	Update(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
	UpdateImgUrl(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
	Delete(ctx context.Context, tx *sql.Tx, SKU string)
}
