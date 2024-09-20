package repository

import (
	"context"
	"database/sql"
	"inventory-system-api/helper"
	"inventory-system-api/model/domain"
)

type LogActivityRepositoryImpl struct {
}

func NewLogActivityRepositoryImpl() LogActivityRepository {
	return &LogActivityRepositoryImpl{}
}

func (repository *LogActivityRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, log domain.LogActivity) {
	SQL := "INSERT INTO log_activity(admin_id, message, time ) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, log.AdminId, log.Message, log.Time)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	log.Id = int(id)

	SQL = "SELECT id, admin_id, message, time FROM log_activity WHERE id = ?"
	err = tx.QueryRowContext(ctx, SQL, log.Id).Scan(&log.Id, &log.AdminId, &log.Message, &log.Time)
	helper.PanicError(err)
}

func (repository *LogActivityRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.LogActivity{
	SQL := "SELECT id, admin_id, message, time FROM log_activity"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()
	
	var logs []domain.LogActivity
	for rows.Next(){
		log := domain.LogActivity{}
		err := rows.Scan(&log.Id, &log.AdminId, &log.Message, &log.Time)
		helper.PanicError(err)

		logs = append(logs, log)
	}

	return logs
}
