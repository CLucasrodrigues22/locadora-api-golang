package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"mime"
	"net/url"
	"os"
	"strings"
)

func GetEnv(variable string, logger *logs.Logger) string {
	err := godotenv.Load()

	if err != nil {
		logger.Errorf("Env file not found: %v", err)
	}

	return os.Getenv(variable)
}

func GenerateFileName(contentType string) string {
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
