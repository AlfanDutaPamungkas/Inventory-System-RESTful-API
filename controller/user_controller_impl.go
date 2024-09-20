package controller

import (
	"context"
	"inventory-system-api/helper"
	"inventory-system-api/model/web"
	"inventory-system-api/service"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type UsersControllerImpl struct {
	UsersService       service.UsersService
	LogActivityService service.LogActivityService
}

func NewUsersControllerImpl(userService service.UsersService, logActivityService service.LogActivityService) UsersController {
	return &UsersControllerImpl{
		UsersService:       userService,
		LogActivityService: logActivityService,
	}
}

func (controller *UsersControllerImpl) CreateAdminCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userCreateReq := web.UsersCreateReq{}
	helper.BodyToReq(r, &userCreateReq)

	response := controller.UsersService.CreateAdminService(r.Context(), userCreateReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	message := "Create admin account with id " + response.Id
	controller.LogActivityService.CreateService(r.Context(), message, response.UpdatedAt)

	helper.WriteToBody(w, webResponse)
}

func (controller *UsersControllerImpl) LoginCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userLoginReq := web.UsersLoginReq{}
	helper.BodyToReq(r, &userLoginReq)

	response := controller.UsersService.LoginService(r.Context(), userLoginReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *UsersControllerImpl) ProfileCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userData := r.Context().Value("userData").(*web.TokenClaims)
	response := controller.UsersService.ProfileService(r.Context(), userData.UserId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *UsersControllerImpl) UpdateAdminAccCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUpdateReq := web.UsersUpdateReq{}
	helper.BodyToReq(r, &userUpdateReq)

	userUpdateReq.Id = p.ByName("id")

	response := controller.UsersService.UpdateAdminAccService(r.Context(), userUpdateReq, userUpdateReq.Id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	message := "Update admin account with id " + response.Id
	controller.LogActivityService.CreateService(r.Context(), message, response.UpdatedAt)

	helper.WriteToBody(w, webResponse)
}

func (controller *UsersControllerImpl) UpdateProfileCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUpdateReq := web.UsersUpdateReq{}
	helper.BodyToReq(r, &userUpdateReq)

	userUpdateReq.Id = r.Context().Value("userData").(*web.TokenClaims).UserId

	response := controller.UsersService.UpdateAdminAccService(r.Context(), userUpdateReq, userUpdateReq.Id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	message := "Create profile account"
	controller.LogActivityService.CreateService(r.Context(), message, response.UpdatedAt)

	helper.WriteToBody(w, webResponse)
}

func (controller *UsersControllerImpl) FindAllAdminAccCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	status := r.URL.Query().Get("status")
	name := r.URL.Query().Get("name")

	queries := map[string]string{
		"status": status,
		"name":   name,
	}

	ctx := context.WithValue(r.Context(), "queries", queries)
	r = r.WithContext(ctx)

	response := controller.UsersService.FindAllAdminAccService(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *UsersControllerImpl) DeactiveAdminCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response := controller.UsersService.DeactiveAdminService(r.Context(), p.ByName("id"))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	message := "Deactive admin account with id " + response.Id
	controller.LogActivityService.CreateService(r.Context(), message, response.UpdatedAt)

	helper.WriteToBody(w, webResponse)
}

func (controller *UsersControllerImpl) FindAdminByIdCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userData := r.Context().Value("userData").(*web.TokenClaims)
	if userData.ID == p.ByName("id") {
		http.Redirect(w, r, "/api/auth/profile", http.StatusMovedPermanently)
		return
	}

	response := controller.UsersService.ProfileService(r.Context(), p.ByName("id"))

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToBody(w, webResponse)
}

func (controller *UsersControllerImpl) ChangePasswordCtrl(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUpdatePassReq := web.UserUpdatePasswordReq{}
	helper.BodyToReq(r, &userUpdatePassReq)
	response := controller.UsersService.ChangePasswordService(r.Context(), userUpdatePassReq)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	message := "Admin with id " + r.Context().Value("userData").(*web.TokenClaims).UserId + "change passsword account"
	controller.LogActivityService.CreateService(r.Context(), message, time.Now())

	helper.WriteToBody(w, webResponse)
}
