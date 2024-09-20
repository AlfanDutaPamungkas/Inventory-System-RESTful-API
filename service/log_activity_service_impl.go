package service

import (
	"context"
	"database/sql"
	"inventory-system-api/helper"
	"inventory-system-api/model/domain"
	"inventory-system-api/model/web"
	"inventory-system-api/repository"
	"time"
)

type LogActivityServiceImpl struct {
	DB            *sql.DB
	LogRepository repository.LogActivityRepository
}

func NewLogActivityServiceImpl(DB *sql.DB, logRepository repository.LogActivityRepository) LogActivityService {
	return &LogActivityServiceImpl{
		DB:            DB,
		LogRepository: logRepository,
	}
}

func (service *LogActivityServiceImpl) CreateService(ctx context.Context, message string, time time.Time) {
	id := ctx.Value("userData").(*web.TokenClaims).UserId

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	log := domain.LogActivity{
		AdminId: id,
		Message: message,
		Time: time,
	}

	service.LogRepository.Create(ctx, tx, log)
}

func (service *LogActivityServiceImpl) FindAllService(ctx context.Context) []domain.LogActivity{
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	logs := service.LogRepository.FindAll(ctx, tx)
	return logs
}
