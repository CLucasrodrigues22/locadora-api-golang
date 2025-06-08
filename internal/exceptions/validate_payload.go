package exceptions

import "github.com/gin-gonic/gin"

func ValidatePayload(ctx *gin.Context, err error, status int, message string) bool {
	if err != nil {
		ctx.JSON(status, gin.H{"error": message})
		return true
	}
	return false
}
