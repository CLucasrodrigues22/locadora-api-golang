package configs

import "github.com/CLucasrodrigues22/api-locadora/internal/configs/logs"

var (
	logger *logs.Logger
)

func GetLogger(pfx string) *logs.Logger {
	logger := logs.NewLogger(pfx)

	return logger
}
