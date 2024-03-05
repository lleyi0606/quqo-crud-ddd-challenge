package db

import (
	"errors"
	loggerentity "products-crud/domain/entity/logger_entity"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

func NewLoggerDB() (*loggerentity.Logger, error) {

	// honeycomb for tracing
	tracer := otel.Tracer("backend") // honeycomb.io

	if tracer == nil {
		zap.S().Errorw("Tracer initialise error")
		return nil, errors.New("error tracer connection")
	}

	// zap for logging
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	undo := zap.ReplaceGlobals(zapLogger)
	defer undo()

	logger := loggerentity.Logger{
		HoneycombTracer: tracer,
		ZapLogger:       zapLogger,
	}

	return &logger, nil

}
