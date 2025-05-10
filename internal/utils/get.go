package utils

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(variable string) string {
	logger := GetLogger("Get .Env variables")

	err := godotenv.Load()

	if err != nil {
		logger.Errorf("Env file not found: %v", err)
	}

	return os.Getenv(variable)
}

func GetLogger(pfx string) *logs.Logger {
	logger := logs.NewLogger(pfx)

	return logger
}
