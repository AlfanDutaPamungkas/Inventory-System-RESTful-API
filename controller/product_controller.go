package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductsController interface {
	CreateController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAllController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindBySKUController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	StockOutController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	StockInController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateImgUrlController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	NullifyExpiredDateController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeleteController(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
