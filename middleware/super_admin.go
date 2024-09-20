package middleware

import (
	"inventory-system-api/helper"
	"inventory-system-api/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SuperAdminMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		userData := r.Context().Value("userData").(*web.TokenClaims)
		if userData.Role != "super admin"{
			forbidden(w)
			return
		}
		
		next(w, r, p)
	}
}

func forbidden(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code: http.StatusForbidden,
		Status: "FORBIDDEN",
	}

	helper.WriteToBody(w, webResponse)
}

