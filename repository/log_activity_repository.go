package repository

import (
	"context"
	"database/sql"
	"inventory-system-api/model/domain"
)

type LogActivityRepository interface {
	Create(ctx context.Context, tx *sql.Tx, log domain.LogActivity)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.LogActivity
}
