package configs

import (
	"fmt"
	"github.com/CLucasrodrigues22/api-locadora/internal/schemas"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbHost     = utils.GetEnv("DB_HOST")
	dbPort     = utils.GetEnv("DB_PORT")
	dbName     = utils.GetEnv("DB_NAME")
	dbUser     = utils.GetEnv("DB_USER")
	dbPassword = utils.GetEnv("DB_PASSWORD")
)

func InitializeDB() (*gorm.DB, error) {
	logger := utils.GetLogger("InitializeDB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("Database connect error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Brand{})

	if err != nil {
		logger.Errorf("mysql migrate err: %v", err)
		return nil, err
	}

	return db, nil
}
