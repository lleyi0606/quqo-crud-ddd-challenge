package logger

import (
	"context"
	"log"
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
	// p            *base.Persistence
	c            *gin.Context
	span         trace.Span
	loggers      []logger_repository.LoggerRepository
	Otel_context *context.Context
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
			// hcRepo = honeycomb.NewHoneycombRepository()
			// loggers = append(loggers, hcRepo)
		}
	}

	return &LoggerRepo{
		// p:            &base.Persistence{},
		c:            &gin.Context{},
		span:         nil,
		loggers:      loggers,
		Otel_context: &hcRepo.Otel_context,
	}
}

func (l *LoggerRepo) Start(c *gin.Context, functionPath string, fields map[string]interface{}, options ...Option) func() {

	l.c = c

	var endFunc func()
	for _, logger := range l.loggers {
		if _, ok := logger.(*honeycomb.HoneycombRepo); ok {
			endFunc = logger.Start(c, functionPath, fields)
		} else {
			logger.Start(c, functionPath, fields)
		}
	}

	for _, opt := range options {
		opt(l)
	}

	return endFunc

}

func (l *LoggerRepo) Debug(msg string, fields map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Debug(msg, fields)
	}
}

func (l *LoggerRepo) Info(msg string, fields map[string]interface{}) {

	log.Print("!!! info called")

	for _, logger := range l.loggers {
		log.Print("logger in Info")
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
		// c.c.Set("otel-context", c.Otel_context)
	}
}

func (l *LoggerRepo) SetContextFromSpan() {
	newCtx := trace.ContextWithSpan(l.c.Request.Context(), l.span)
	l.c.Set("otel-context", newCtx)
}

func (l *LoggerRepo) End() {

	for _, logger := range l.loggers {
		logger.End()
	}

	// l.c.Set("otel-context", l.Otel_context)
	// l.span.End()
	// log.Print("otel context set in End()")

	// var ctx context.Context
	// if storedCtx, exists := l.c.Get("otel-context"); exists {
	// 	log.Println("otel-context FOUND in End")
	// 	ctx = storedCtx.(context.Context)
	// } else {
	// 	log.Println("otel-context NOT FOUND in End")
	// 	ctx = l.c.Request.Context()
	// }

	// span := trace.SpanFromContext(ctx)
	// defer span.End()

}
