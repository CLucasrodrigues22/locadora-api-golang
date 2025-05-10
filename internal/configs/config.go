package configs

import (
	"fmt"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"gorm.io/gorm"
)

var (
	logger *logs.Logger
	db     *gorm.DB
)

func Init() error {
	var err error

	// Initialize Database
	db, err = InitializeDB()
	if err != nil {
		return fmt.Errorf("init mysql err: %v", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
