package repository

import (
	"context"
	"database/sql"
	"inventory-system-api/model/domain"
)

type UsersRepository interface {
	CreateAdmin(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Users, error)
	FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Users, error)
	UpdateAdminAcc(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Users
	DeactiveAdmin(ctx context.Context, tx *sql.Tx, id string) domain.Users
	ChangePassword(ctx context.Context, tx *sql.Tx, user domain.Users)
}
