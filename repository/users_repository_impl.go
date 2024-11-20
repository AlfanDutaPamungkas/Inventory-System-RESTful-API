package repository

import (
	"context"
	"database/sql"
	"errors"
	"inventory-system-api/helper"
	"inventory-system-api/model/domain"
)

type UsersRepositoryImpl struct{}

func NewUserRepositoryImpl() UsersRepository {
	return &UsersRepositoryImpl{}
}

func (repository *UsersRepositoryImpl) CreateAdmin(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "INSERT INTO users(id, name, email, password) values (?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Name, user.Email, user.Password)
	helper.PanicError(err)

	SQL = "SELECT id, name, email, role, status, created_at, updated_at FROM users WHERE id = ?"
	err = tx.QueryRowContext(ctx, SQL, user.Id).Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	helper.PanicError(err)

	return user
}

func (repository *UsersRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.Users, error) {
	SQL := "SELECT id, name, email, password, role, status, created_at, updated_at FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicError(err)
	defer rows.Close()

	user := domain.Users{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicError(err)
		return user, nil
	} else {
		return user, errors.New("users not found")
	}
}

func (repository *UsersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (domain.Users, error) {
	SQL := "SELECT id, name, email, role, status, created_at, updated_at FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicError(err)
	defer rows.Close()

	user := domain.Users{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicError(err)
		return user, nil
	} else {
		return user, errors.New("users not found")
	}
}

func (repository *UsersRepositoryImpl) UpdateAdminAcc(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Id)
	helper.PanicError(err)

	SQL = "SELECT id, name, email, role, status, created_at, updated_at FROM users WHERE id = ?"
	err = tx.QueryRowContext(ctx, SQL, user.Id).Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	helper.PanicError(err)

	return user
}

func (repository *UsersRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Users {
	queries := ctx.Value("queries").(map[string]string)
	var rows *sql.Rows
	var err error

	if queries["status"] != "" && queries["name"] != ""{
		SQL := "SELECT id, name, email, role, status, created_at, updated_at FROM users WHERE  MATCH (name) AGAINST (? IN NATURAL LANGUAGE MODE) AND status = ?"
		rows, err = tx.QueryContext(ctx, SQL, queries["name"], queries["status"])
	}else if queries["name"] != "" {
		SQL := "SELECT id, name, email, role, status, created_at, updated_at FROM users WHERE MATCH (name) AGAINST (? IN NATURAL LANGUAGE MODE)"
		rows, err = tx.QueryContext(ctx, SQL, queries["name"])
	} else if queries["status"] != "" {
		SQL := "SELECT id, name, email, role, status, created_at, updated_at FROM users where status = ?"
		rows, err = tx.QueryContext(ctx, SQL, queries["status"])
	} else {
		SQL := "SELECT id, name, email, role, status, created_at, updated_at FROM users"
		rows, err = tx.QueryContext(ctx, SQL)
	}
	helper.PanicError(err)
	defer rows.Close()

	var users []domain.Users
	for rows.Next() {
		user := domain.Users{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicError(err)

		users = append(users, user)
	}

	return users
}

func (repository *UsersRepositoryImpl) DeactiveAdmin(ctx context.Context, tx *sql.Tx, id string) domain.Users {
	SQL := "UPDATE users SET status = 'inactive' WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicError(err)

	user := domain.Users{}

	SQL = "SELECT id, name, email, role, status, created_at, updated_at FROM users WHERE id = ?"
	err = tx.QueryRowContext(ctx, SQL, id).Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	helper.PanicError(err)

	return user
}

func (repository *UsersRepositoryImpl) ChangePassword(ctx context.Context, tx *sql.Tx, user domain.Users) {
	SQL := "UPDATE users SET password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Id)
	helper.PanicError(err)
}
