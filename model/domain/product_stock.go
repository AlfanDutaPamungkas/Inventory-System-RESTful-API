package domain

import "time"

type ProductStock struct {
	Id, Amount  int
	SKU         string
	ExpiredDate *time.Time
	UpdatedAt   time.Time
}
