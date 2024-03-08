package loggerentity

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

type Logger struct {
	HoneycombTracer *sdktrace.TracerProvider
	ZapLogger       *zap.Logger
}

type FunctionInfo struct {
	FunctionName string `gorm:"size:255;"`
	Path         string `gorm:"size:255;"`
	Description  string `gorm:"size:255;"`
	Body         interface{}
}
