package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsersController interface {
	CreateAdminCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	LoginCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	ProfileCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateProfileCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	UpdateAdminAccCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAllAdminAccCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	DeactiveAdminCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAdminByIdCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	ChangePasswordCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
