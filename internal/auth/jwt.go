package auth

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/common"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtKey = []byte(common.GetEnv("JWT_SECRET"))

func GenerateJWT(email string, name string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := &schemas.Claims{
		Email: email,
		Name:  name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "issuer-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
