package service

import (
	"context"
	"inventory-system-api/model/domain"
	"time"
)

type LogActivityService interface {
	CreateService(ctx context.Context, message string, time time.Time)
	FindAllService(ctx context.Context) []domain.LogActivity
}
