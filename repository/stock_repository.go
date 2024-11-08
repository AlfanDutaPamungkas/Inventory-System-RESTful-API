package repository

import (
	"context"
	"database/sql"
	"inventory-system-api/model/domain"
)

type StockRepository interface {
	StockOut(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
	NullifyExpiredDate(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
}
