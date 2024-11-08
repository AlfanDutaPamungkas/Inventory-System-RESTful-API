package repository

import (
	"context"
	"database/sql"
	"inventory-system-api/helper"
	"inventory-system-api/model/domain"
)

type StockRepositoryImpl struct {
}

func NewStockRepositoryImpl() StockRepository {
	return &StockRepositoryImpl{}
}

func (repository *StockRepositoryImpl) StockOut(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	SQL := "UPDATE product_stock SET stock_amount = stock_amount + ? WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Amount, product.SKU)
	helper.PanicError(err)

	SQL = `SELECT 
				products.SKU,
				products.product_name,
				products.product_brand,
				products.category,
				products.price,
				products.image_url,
				product_stock.stock_amount,
				product_stock.expired_date,
				products.created_at,
				GREATEST(products.updated_at, product_stock.updated_at) AS latest_update_at
			FROM
			products
				JOIN
			product_stock ON (products.SKU = product_stock.SKU)
			WHERE products.SKU = ?`
	err = tx.QueryRowContext(ctx, SQL, product.SKU).Scan(&product.SKU, &product.Name, &product.Brand, &product.Category, &product.Price, &product.ImageUrl, &product.Amount, &product.ExpiredDate, &product.CreatedAt, &product.UpdatedAt)
	helper.PanicError(err)

	return product
}

func (repository *StockRepositoryImpl) NullifyExpiredDate(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	SQL := "UPDATE product_stock SET expired_date = NULL WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, product.SKU)
	helper.PanicError(err)

	SQL = `SELECT 
				products.SKU,
				products.product_name,
				products.product_brand,
				products.category,
				products.price,
				products.image_url,
				product_stock.stock_amount,
				product_stock.expired_date,
				products.created_at,
				GREATEST(products.updated_at, product_stock.updated_at) AS latest_update_at
			FROM
			products
				JOIN
			product_stock ON (products.SKU = product_stock.SKU)
			WHERE products.SKU = ?`
	err = tx.QueryRowContext(ctx, SQL, product.SKU).Scan(&product.SKU, &product.Name, &product.Brand, &product.Category, &product.Price, &product.ImageUrl, &product.Amount, &product.ExpiredDate, &product.CreatedAt, &product.UpdatedAt)
	helper.PanicError(err)

	return product
}

