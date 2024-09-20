package middleware

import (
	"context"
	"inventory-system-api/helper"
	"inventory-system-api/model/web"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		tokenAuth := r.Header.Get("Authorization")
		if tokenAuth == "" {
			unauthorize(w, "unauthorize - please login first")
			return
		}

		parts := strings.SplitN(tokenAuth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			unauthorize(w, "something wrong in your authentication")
			return
		}
		tokenAuth = parts[1]
		jwtTokenSecret := []byte(os.Getenv("JWT_TOKEN_SECRET"))
		claims := &web.TokenClaims{}
		token, err := jwt.ParseWithClaims(tokenAuth, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtTokenSecret, nil
		})

		if err != nil || !token.Valid {
			unauthorize(w, "invalid")
			return
		}

		ctx := context.WithValue(r.Context(), "userData", claims)
		r = r.WithContext(ctx)

		next(w, r, p)
	}
}

func unauthorize(w http.ResponseWriter, v string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Status: "UNAUTHORIZE",
		Code:   http.StatusUnauthorized,
		Data: v,
	}

	helper.WriteToBody(w, webResponse)
}

