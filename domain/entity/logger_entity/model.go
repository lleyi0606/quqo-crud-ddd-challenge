package loggerentity

import "products-crud/domain/repository/logger_repository"

// type Logger struct {
// 	HoneycombTracer *sdktrace.TracerProvider
// 	ZapLogger       *zap.Logger
// }

type Logger logger_repository.LoggerRepository

type FunctionInfo struct {
	FunctionName string `gorm:"size:255;"`
	Path         string `gorm:"size:255;"`
	Description  string `gorm:"size:255;"`
	Body         interface{}
}
