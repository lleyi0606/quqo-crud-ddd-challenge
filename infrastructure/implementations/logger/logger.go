package logger

import (
	"products-crud/domain/repository/logger_repository"
	"products-crud/infrastructure/implementations/logger/honeycomb"
	"products-crud/infrastructure/implementations/logger/zap"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

const (
	Honeycomb = "HONEYCOMB"
	Zap       = "ZAP"
)

type LoggerRepo struct {
	c       *gin.Context
	loggers []logger_repository.LoggerRepository
}

type Option func(*LoggerRepo)

func NewLoggerRepositories(providers []string) *LoggerRepo {

	var loggers []logger_repository.LoggerRepository

	var hcRepo *honeycomb.HoneycombRepo
	for _, provider := range providers {
		switch strings.ToUpper(provider) {
		case Zap:
			loggers = append(loggers, zap.NewZapRepository())
		case Honeycomb:
			hcRepo = honeycomb.NewHoneycombRepository()
			loggers = append(loggers, hcRepo)
		default:
			hcRepo = honeycomb.NewHoneycombRepository()
			loggers = append(loggers, hcRepo)
		}
	}

	return &LoggerRepo{
		c:       &gin.Context{},
		loggers: loggers,
	}
}

func (l *LoggerRepo) Start(c *gin.Context, functionPath string, fields map[string]interface{}, options ...Option) trace.Span {

	l.c = c

	var span trace.Span
	for _, logger := range l.loggers {
		if _, ok := logger.(*honeycomb.HoneycombRepo); ok {
			span = logger.Start(c, functionPath, fields)
		} else {
			logger.Start(c, functionPath, fields)
		}
	}

	for _, opt := range options {
		opt(l)
	}

	return span
}

func (l *LoggerRepo) Debug(msg string, fields map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Debug(msg, fields)
	}
}

func (l *LoggerRepo) Info(msg string, fields map[string]interface{}, options ...Option) {
	for _, opt := range options {
		opt(l)
	}

	for _, logger := range l.loggers {
		logger.Info(msg, fields)
	}
}

func (l *LoggerRepo) Warn(msg string, fields map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Warn(msg, fields)
	}
}

func (l *LoggerRepo) Error(msg string, fields map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Error(msg, fields)
	}
}

func (l *LoggerRepo) Fatal(msg string, fields map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Fatal(msg, fields)
	}
}

func SetNewOtelContext() Option {
	return func(c *LoggerRepo) {
		for _, logger := range c.loggers {
			if honeycombRepo, ok := logger.(*honeycomb.HoneycombRepo); ok {
				honeycombRepo.SetNewOtelContext()
			}
		}
	}
}

func WithSpan(span trace.Span) Option {
	return func(c *LoggerRepo) {
		for _, logger := range c.loggers {
			if honeycombRepo, ok := logger.(*honeycomb.HoneycombRepo); ok {
				honeycombRepo.UseGivenSpan(span)
			}
		}
	}
}

func (l *LoggerRepo) SetContextFromSpan(span trace.Span) {
	newCtx := trace.ContextWithSpan(l.c.Request.Context(), span)
	l.c.Set("otel-context", newCtx)
}

// func (l *LoggerRepo) End() {
// 	for _, logger := range l.loggers {
// 		logger.End()
// 	}
// }
