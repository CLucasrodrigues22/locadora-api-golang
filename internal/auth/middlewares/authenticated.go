package middlewares

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/auth"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			ctx.Abort()
			return
		}

		tokenJWT := parts[1]

		claims := &schemas.Claims{}
		token, err := jwt.ParseWithClaims(tokenJWT, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenMalformed
			}
			return auth.JwtKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("name", claims.Name)

		ctx.Next()
	}
}
