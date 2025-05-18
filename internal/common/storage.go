package common

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

func SaveFileStorage(content io.Reader, contentType string) (string, error) {
	s3Client := GetStorageConnection()
	bucket := GetEnv("AWS_BUCKET")
	fileName := generateFileName(contentType)
	region := GetEnv("AWS_DEFAULT_REGION")

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
