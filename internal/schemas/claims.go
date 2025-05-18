package schemas

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}
