package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SendError(context *gin.Context, code int, msg string) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, gin.H{
		"massage": msg,
		"error":   code,
	})
}

func SendSuccess(context *gin.Context, op string, data interface{}, code int) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, gin.H{
		"massage": fmt.Sprintf("operation %s successfully", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
