package repository

import (
	"context"
	"database/sql"
	"errors"
	"inventory-system-api/helper"
	"inventory-system-api/model/domain"
)

type ProductsRepositoryImpl struct{}

func NewProductsRepositoryImpl() ProductsRepository {
	return &ProductsRepositoryImpl{}
}

func (repository *ProductsRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	SQL := "INSERT INTO products(SKU, product_name, product_brand, category, price, image_url) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, product.SKU, product.Name, product.Brand, product.Category, product.Price, product.ImageUrl)
	helper.PanicError(err)

	SQL = "INSERT INTO product_stock(stock_amount, SKU, expired_date) VALUES (?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, product.Amount, product.SKU, product.ExpiredDate)
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
			products.updated_at
		   FROM
		   products
			JOIN
		   product_stock ON (products.SKU = product_stock.SKU)
		   WHERE products.SKU = ?`

	err = tx.QueryRowContext(ctx, SQL, product.SKU).Scan(&product.SKU, &product.Name, &product.Brand, &product.Category, &product.Price, &product.ImageUrl, &product.Amount, &product.ExpiredDate, &product.CreatedAt, &product.UpdatedAt)
	helper.PanicError(err)

	return product
}

func (repository *ProductsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Products {
	name := ctx.Value("query_name").(string)
	var rows *sql.Rows
	var err error

	if name == "" {
		SQL := `
			SELECT 
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
				product_stock ON (products.SKU = product_stock.SKU)`
		rows, err = tx.QueryContext(ctx, SQL)
	} else {
		SQL := `
				SELECT 
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
				WHERE
					MATCH (products.product_name) AGAINST (? IN NATURAL LANGUAGE MODE);
				`
		rows, err = tx.QueryContext(ctx, SQL, name)
	}

	helper.PanicError(err)
	defer rows.Close()

	var products []domain.Products
	for rows.Next() {
		product := domain.Products{}
		err := rows.Scan(&product.SKU, &product.Name, &product.Brand, &product.Category, &product.Price, &product.ImageUrl, &product.Amount, &product.ExpiredDate, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicError(err)

		products = append(products, product)
	}
	return products
}

func (repository *ProductsRepositoryImpl) FindBySKU(ctx context.Context, tx *sql.Tx, SKU string) (domain.Products, error) {
	SQL := `SELECT 
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
	rows, err := tx.QueryContext(ctx, SQL, SKU)
	helper.PanicError(err)
	defer rows.Close()

	product := domain.Products{}

	if rows.Next() {
		err = rows.Scan(&product.SKU, &product.Name, &product.Brand, &product.Category, &product.Price, &product.ImageUrl, &product.Amount, &product.ExpiredDate, &product.CreatedAt, &product.UpdatedAt)
		helper.PanicError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}
}

func (repository *ProductsRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	SQL := "UPDATE products SET product_name = ?, product_brand = ?, category = ?, price = ? WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Brand, product.Category, product.Price, product.SKU)
	helper.PanicError(err)

	SQL = "UPDATE product_stock set stock_amount = ?, expired_date = ? WHERE SKU = ?"
	_, err = tx.ExecContext(ctx, SQL, product.Amount, product.ExpiredDate, product.SKU)
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

func (repository *ProductsRepositoryImpl) UpdateImgUrl(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	SQL := "UPDATE products SET image_url = ? WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, product.ImageUrl, product.SKU)
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

func (repository *ProductsRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, SKU string){
	SQL := "DELETE FROM products WHERE SKU = ?"
	_, err := tx.ExecContext(ctx, SQL, SKU)
	helper.PanicError(err)
}