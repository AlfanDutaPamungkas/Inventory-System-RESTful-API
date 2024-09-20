package helper

import (
	"inventory-system-api/model/web"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(request web.TokenCreateReq, v time.Duration) string {
	jwtTokenSecret := []byte(os.Getenv("JWT_TOKEN_SECRET"))

	expiredTime := time.Now().Add(time.Minute * v)
	claims := &web.TokenClaims{
		UserId: request.UserId,
		Role:   request.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtTokenSecret)
	PanicError(err)

	return tokenStr
}
