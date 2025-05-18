package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"io"
	"mime"
	"net/url"
	"os"
	"strings"
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

func StorageFile(content io.Reader, contentType string) (string, error) {
	fileName := generateFileName(contentType)
	s3Client := GetStorageConnection()
	region := GetEnv("AWS_DEFAULT_REGION")
	bucket := GetEnv("AWS_BUCKET")

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
	bucket := GetEnv("AWS_BUCKET")
	client := GetStorageConnection()

	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	if err != nil {
		logger := GetLogger("Delete File")
		logger.Errorf("Failed to delete file: %v", err)
		return err
	}

	return nil
}

func generateFileName(contentType string) string {
	u := uuid.New().String()
	hash := md5.Sum([]byte(u))
	hashed := hex.EncodeToString(hash[:])

	exts, _ := mime.ExtensionsByType(contentType)
	ext := ".bin"
	if len(exts) > 0 {
		ext = exts[0]
	}

	ext = strings.TrimPrefix(ext, ".")

	fileName := fmt.Sprintf("%s.%s", hashed, ext)
	return fileName
}

func ExtractKeyFromURL(fileURL string) string {
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return fileURL
	}

	return strings.TrimPrefix(parsedURL.Path, "/")
}
