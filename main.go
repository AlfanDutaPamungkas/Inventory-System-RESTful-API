package main

import (
	"inventory-system-api/helper"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr: ":3000",
		Handler: router,
	}
}

func main() {
	envErr := godotenv.Load(".env")
	helper.PanicError(envErr)

	server := InitializedServer()

	log.Println("Server is running on port 3000")
	err := server.ListenAndServe()
	helper.PanicError(err)
}
