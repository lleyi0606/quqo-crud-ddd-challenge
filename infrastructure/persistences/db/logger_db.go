package db

import (
	"os"
	"products-crud/infrastructure/implementations/logger"
	"strings"
)

func NewLoggerDB() (*logger.LoggerRepo, error) {

	loggerTechnology := os.Getenv("LOGGER_DEFAULT_TECHNOLOGY")
	loggersStrArray := strings.Split(loggerTechnology, ",")
	logger := logger.NewLoggerRepositories(loggersStrArray)

	return logger, nil

}
