package main

import (
	"inventory-system-api/app"
	"inventory-system-api/controller"
	"inventory-system-api/helper"
	"inventory-system-api/repository"
	"inventory-system-api/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	helper.PanicError(envErr)

	db := app.Database()
	validate := validator.New()
	cld := app.NewCloudinary()

	logRepp := repository.NewLogActivityRepositoryImpl()
	logService := service.NewLogActivityServiceImpl(db, logRepp)
	LogController := controller.NewLogControllerImpl(logService)

	usersRepo := repository.NewUserRepositoryImpl()
	usersService := service.NewUsersServiceImpl(db, validate, usersRepo)
	usersController := controller.NewUsersControllerImpl(usersService, logService)

	productsRepo := repository.NewProductsRepositoryImpl()
	stockRepo := repository.NewStockRepositoryImpl()
	productsService := service.NewProductServiceImpl(db, productsRepo, stockRepo, cld)
	productsController := controller.NewProductsControllerImpl(productsService, logService)

	router := app.NewRouter(usersController, productsController, LogController)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("Server is running on port 3000")
	err := server.ListenAndServe()
	helper.PanicError(err)
}
