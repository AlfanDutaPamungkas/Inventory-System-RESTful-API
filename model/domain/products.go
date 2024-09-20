package domain

import "time"

type Products struct {
	SKU, Name, Brand, Category, ImageUrl string
	Price                                int
	CreatedAt, UpdatedAt                 time.Time
}
