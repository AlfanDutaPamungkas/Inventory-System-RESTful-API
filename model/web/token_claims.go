package web

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserId string
	Role   string
	jwt.RegisteredClaims
}
