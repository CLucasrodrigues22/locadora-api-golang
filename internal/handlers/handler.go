package handlers

import (
	"github.com/CLucasrodrigues22/api-locadora/internal/configs"
	"github.com/CLucasrodrigues22/api-locadora/internal/logs"
	"github.com/CLucasrodrigues22/api-locadora/internal/utils"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB
	logger *logs.Logger
)

func InitHandler() {
	logger = utils.GetLogger("handler")
	Db = configs.GetDB()
}
