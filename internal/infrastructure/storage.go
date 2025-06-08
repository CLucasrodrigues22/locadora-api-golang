package infrastructure

import (
	"context"
	"fmt"
	"github.com/CLucasrodrigues22/api-locadora/internal/bootstrap"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

var (
	logger *logs.Logger
)

func SaveFileStorage(content io.Reader, contentType string) (string, error) {
	s3Client := bootstrap.GetStorageConnection()
	bucket := utils.GetEnv("AWS_BUCKET", logger)
	fileName := utils.GenerateFileName(contentType)
	region := utils.GetEnv("AWS_DEFAULT_REGION", logger)

	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileName),
		Body:        content,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	urlPath := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, fileName)

	return urlPath, nil
}

func DeleteFileStorage(key string) error {
	bucket := utils.GetEnv("AWS_BUCKET", logger)
	client := bootstrap.GetStorageConnection()

	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	if err != nil {
		logger := bootstrap.GetLogger("Delete File")
		logger.Errorf("Failed to delete file: %v", err)
		return err
	}

	return nil
}
