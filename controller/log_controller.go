package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LogController interface {
	FindAllCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
