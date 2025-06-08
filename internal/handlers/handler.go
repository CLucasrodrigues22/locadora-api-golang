package handlers

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/bootstrap"
	"github.com/CLucasrodrigues22/api-locadora/internal/configs"
	"github.com/CLucasrodrigues22/api-locadora/internal/contracts"
	"github.com/CLucasrodrigues22/api-locadora/internal/infrastructure"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
)

var (
	Db     *gorm.DB
	logger *logs.Logger
)

func InitHandler() {
	Db = configs.GetDB()
	logger = bootstrap.GetLogger("handler")
}

func SaveFile(ctx *gin.Context, file *multipart.FileHeader) (string, error) {
	if file == nil {
		return "", nil
	}

	return processFileUpload(ctx, file)
}

func UpdateFile(ctx *gin.Context, file *multipart.FileHeader, record contracts.HasImage) (string, error) {
	if file == nil {
		return record.GetImagePath(), nil
	}

	oldIcon := record.GetImagePath()
	if oldIcon != "" {
		key := utils.ExtractKeyFromURL(oldIcon)
		err := infrastructure.DeleteFileStorage(key)
		if err != nil {
			logger.Errorf("Failed to delete old file: %v", err)
			SendError(ctx, http.StatusInternalServerError, "Failed to delete old file")
			return "", err
		}
	}

	fileURL, err := processFileUpload(ctx, file)
	if err != nil {
		return "", err
	}

	record.SetImagePath(fileURL)
	return fileURL, nil
}

func DeleteFile(ctx *gin.Context, record contracts.HasImage) error {
	icon := record.GetImagePath()
	if icon == "" {
		return nil
	}

	key := utils.ExtractKeyFromURL(icon)

	if err := infrastructure.DeleteFileStorage(key); err != nil {
		logger.Errorf("Failed to delete file: %v", err)
		SendError(ctx, http.StatusInternalServerError, "Failed to delete file")
		return err
	}

	record.SetImagePath("")
	return nil
}

func processFileUpload(ctx *gin.Context, file *multipart.FileHeader) (string, error) {
	fileReader, err := file.Open()
	if err != nil {
		SendError(ctx, http.StatusBadRequest, "Invalid file")
		return "", err
	}
	defer fileReader.Close()

	contentType := file.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	fileURL, err := infrastructure.SaveFileStorage(fileReader, contentType)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, "Failed to upload file")
		return "", err
	}

	return fileURL, nil
}
