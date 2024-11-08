package repository

import (
	"context"
	"database/sql"
	"inventory-system-api/model/domain"
)

type StockRepository interface {
	Create(ctx context.Context, tx *sql.Tx, stock domain.ProductStock) domain.ProductStock
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductStock
	FindBySKU(ctx context.Context, tx *sql.Tx, SKU string) (domain.ProductStock, error)
	Update(ctx context.Context, tx *sql.Tx, stock domain.ProductStock) domain.ProductStock
	StockOut(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
	NullifyExpiredDate(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
	Delete(ctx context.Context, tx *sql.Tx, SKU string)
}
