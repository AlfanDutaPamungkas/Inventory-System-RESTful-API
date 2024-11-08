package domain

import "time"

type Products struct {
	SKU, Name, Brand, Category, ImageUrl string
	Price, Amount                        int
	ExpiredDate                          *time.Time
	CreatedAt, UpdatedAt                 time.Time
}
