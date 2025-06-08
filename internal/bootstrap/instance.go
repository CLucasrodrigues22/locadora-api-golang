package bootstrap

import (
	"context"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	logger *logs.Logger
)

func GetLogger(pfx string) *logs.Logger {
	logger = logs.NewLogger(pfx)

	return logger
}

func GetStorageConnection() *s3.Client {
	accessKeyId := utils.GetEnv("AWS_ACCESS_KEY_ID", logger)
	secretAccessKey := utils.GetEnv("AWS_SECRET_ACCESS_KEY", logger)

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
