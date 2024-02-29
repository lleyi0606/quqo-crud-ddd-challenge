package loggerentity

import (
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Logger struct {
	HoneycombTracer *trace.Tracer
	ZapLogger       *zap.Logger
}

type FunctionInfo struct {
	FunctionName string `gorm:"size:255;"`
	Path         string `gorm:"size:255;"`
	Description  string `gorm:"size:255;"`
	Body         interface{}
}
