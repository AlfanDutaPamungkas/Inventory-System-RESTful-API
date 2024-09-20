package repository

import (
	"context"
	"database/sql"
	"errors"
	"inventory-system-api/helper"
	"inventory-system-api/model/domain"
)

type StockRepositoryImpl struct {
}

func NewStockRepositoryImpl() StockRepository {
	return &StockRepositoryImpl{}
}

func (repository *StockRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, stock domain.ProductStock) domain.ProductStock {
	SQL := "INSERT INTO product_stock(stock_amount, SKU, expired_date) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, stock.Amount, stock.SKU, stock.ExpiredDate)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	stock.Id = int(id)

	SQL = "SELECT id, stock_amount, SKU, expired_date FROM product_stock WHERE id = ?"
	err = tx.QueryRowContext(ctx, SQL, stock.Id).Scan(&stock.Id, &stock.Amount, &stock.SKU, &stock.ExpiredDate)
	helper.PanicError(err)

	return stock
}

func (repository *StockRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductStock {
	SQL := "SELECT id, SKU, stock_amount, expired_date, updated_at FROM product_stock"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()

	var stocks []domain.ProductStock
	for rows.Next() {
		stock := domain.ProductStock{}
		err := rows.Scan(&stock.Id, &stock.SKU, &stock.Amount, &stock.ExpiredDate, &stock.UpdatedAt)
		helper.PanicError(err)

		stocks = append(stocks, stock)
	}

	return stocks
}

func (repository *StockRepositoryImpl) FindBySKU(ctx context.Context, tx *sql.Tx, SKU string) (domain.ProductStock, error) {
	SQL := "SELECT id, stock_amount, SKU, expired_date, updated_at FROM product_stock WHERE SKU = ?"
	rows, err := tx.QueryContext(ctx, SQL, SKU)
	helper.PanicError(err)
	defer rows.Close()

	stock := domain.ProductStock{}

	if rows.Next() {
		err = rows.Scan(&stock.Id, &stock.Amount, &stock.SKU, &stock.ExpiredDate, &stock.UpdatedAt)
		helper.PanicError(err)
		return stock, nil
	} else {
		return stock, errors.New("product stock not found")
	}
}

func (repository *StockRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, stock domain.ProductStock) domain.ProductStock {
	SQL := "UPDATE product_stock set stock_amount = ?, expired_date = ? WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, stock.Amount, stock.ExpiredDate, stock.SKU)
	helper.PanicError(err)

	SQL = "SELECT id, stock_amount, SKU, expired_date FROM product_stock WHERE id = ?"
	err = tx.QueryRowContext(ctx, SQL, stock.Id).Scan(&stock.Id, &stock.Amount, &stock.SKU, &stock.ExpiredDate)
	helper.PanicError(err)

	return stock
}

func (repository *StockRepositoryImpl) StockOut(ctx context.Context, tx *sql.Tx, stock domain.ProductStock) domain.ProductStock {
	SQL := "UPDATE product_stock SET stock_amount = stock_amount + ? WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, stock.Amount, stock.SKU)
	helper.PanicError(err)

	SQL = "SELECT id, stock_amount, SKU, expired_date FROM product_stock WHERE id = ?"
	err = tx.QueryRowContext(ctx, SQL, stock.Id).Scan(&stock.Id, &stock.Amount, &stock.SKU, &stock.ExpiredDate)
	helper.PanicError(err)

	return stock
}

func (repository *StockRepositoryImpl) NullifyExpiredDate(ctx context.Context, tx *sql.Tx, stock domain.ProductStock) domain.ProductStock {
	SQL := "UPDATE product_stock SET expired_date = NULL WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, stock.SKU)
	helper.PanicError(err)

	SQL = "SELECT id, stock_amount, SKU, expired_date FROM product_stock WHERE id = ?"
	err = tx.QueryRowContext(ctx, SQL, stock.Id).Scan(&stock.Id, &stock.Amount, &stock.SKU, &stock.ExpiredDate)
	helper.PanicError(err)

	return stock
}

func (repository *StockRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, SKU string){
	SQL := "DELETE FROM product_stock WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, SKU)
	helper.PanicError(err)
}

