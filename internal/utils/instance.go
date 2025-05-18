package utils

import (
	"context"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

func GetStorageConnection() *s3.Client {
	accessKeyId := GetEnv("AWS_ACCESS_KEY_ID")
	secretAccessKey := GetEnv("AWS_SECRET_ACCESS_KEY")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, "")),
	)

	if err != nil {
		logger := GetLogger("Get Storage Connection")
		logger.Errorf("Unable to load SDK config to storage, %v", err)
	}

	client := s3.NewFromConfig(cfg)
	return client
}
