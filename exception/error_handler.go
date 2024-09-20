package exception

import (
	"inventory-system-api/helper"
	"inventory-system-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if unauthorize(w, r, err) {
		return
	}

	if forbidden(w, r, err){
		return
	}

	if badReqErr(w,r, err){
		return
	}

	if validationErr(w, r, err) {
		return
	}

	if notFound(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func unauthorize(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		response := web.WebResponse{
			Status: "UNAUTHORIZE",
			Code:   http.StatusUnauthorized,
			Data:   exception.Error,
		}

		helper.WriteToBody(w, response)
		return true
	} else {
		return false
	}
}

func forbidden(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(ForbiddenErr)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		response := web.WebResponse{
			Status: "FORBIDDEN",
			Code:   http.StatusForbidden,
			Data:   exception.Error,
		}

		helper.WriteToBody(w, response)
		return true
	} else {
		return false
	}
}

func badReqErr(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(BadReqErr)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func validationErr(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func notFound(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundErr)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := web.WebResponse{
			Status: "NOT FOUND",
			Code:   http.StatusNotFound,
			Data:   exception.Error,
		}

		helper.WriteToBody(w, response)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}

	helper.WriteToBody(w, response)
}
