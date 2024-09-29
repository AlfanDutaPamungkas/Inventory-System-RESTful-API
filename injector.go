//go:build wireinject
// +build wireinject

package main

import (
	"inventory-system-api/app"
	"inventory-system-api/controller"
	"inventory-system-api/repository"
	"inventory-system-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var logSet = wire.NewSet(
	repository.NewLogActivityRepositoryImpl,
	service.NewLogActivityServiceImpl,
	controller.NewLogControllerImpl,
)

var userSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	service.NewUsersServiceImpl,
	controller.NewUsersControllerImpl,
)

var productSet = wire.NewSet(
	repository.NewProductsRepositoryImpl,
	repository.NewStockRepositoryImpl,
	service.NewProductServiceImpl,
	controller.NewProductsControllerImpl,
)

func ProvideValidator() *validator.Validate {
    return validator.New()
}

func InitializedServer() *http.Server {
	wire.Build(
		app.Database,
		ProvideValidator,
		app.NewCloudinary,
		logSet,
		userSet,
		productSet,
		app.NewRouter,
		NewServer,
	)
	return nil
}
