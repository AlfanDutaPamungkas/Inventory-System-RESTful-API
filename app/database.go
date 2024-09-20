package app

import (
	"database/sql"
	"inventory-system-api/helper"
	"os"
	"time"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_URL"))
	helper.PanicError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
