package logger

import (
	"context"
	"log"
	loggerentity "products-crud/domain/entity/logger_entity"
	"products-crud/domain/repository/logger_repository"
	"products-crud/infrastructure/implementations/logger/honeycomb"
	"products-crud/infrastructure/implementations/logger/zap"
	base "products-crud/infrastructure/persistences"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

const (
	Honeycomb = "HONEYCOMB"
	Zap       = "ZAP"
)

type loggerRepo struct {
	p            *base.Persistence
	c            *gin.Context
	span         trace.Span
	loggers      []logger_repository.LoggerRepository
	Otel_context *context.Context
}

type Option func(*loggerRepo)

// func NewLoggerRepositoriesOld(p *base.Persistence, c *gin.Context, info loggerentity.FunctionInfo, providers ...string) loggerRepo {

// 	_, callerInfo, _, _ := runtime.Caller(1)
// 	log.Printf("!!! new logger repo called, %s", callerInfo)
// 	var loggers []logger_repository.LoggerRepository

// 	var hcRepo *honeycomb.HoneycombRepo
// 	for _, provider := range providers {
// 		switch strings.ToUpper(provider) {
// 		case Zap:
// 			loggers = append(loggers, zap.NewZapRepository(p, c))
// 		case Honeycomb:
// 			hcRepo = honeycomb.NewHoneycombRepository(p, c, info)
// 			loggers = append(loggers, hcRepo)
// 		default:
// 			hcRepo = honeycomb.NewHoneycombRepository(p, c, info)
// 			loggers = append(loggers, hcRepo)
// 		}
// 	}

// 	return loggerRepo{
// 		p:       p,
// 		c:       c,
// 		span:    hcRepo.Span,
// 		loggers: loggers,
// 	}
// }

// func NewLoggerRepositoriesOld(p *base.Persistence, c *gin.Context, info loggerentity.FunctionInfo, providers ...string) []logger_repository.LoggerRepository {
func NewLoggerRepositories(p *base.Persistence, c *gin.Context, info loggerentity.FunctionInfo, providers []string, options ...Option) (loggerRepo, func()) {

	// _, callerInfo, _, _ := runtime.Caller(1)
	// log.Printf("!!! new logger repo called, %s", callerInfo)
	var loggers []logger_repository.LoggerRepository

	var hcRepo *honeycomb.HoneycombRepo
	for _, provider := range providers {
		switch strings.ToUpper(provider) {
		case Zap:
			loggers = append(loggers, zap.NewZapRepository(p, c))
		case Honeycomb:
			hcRepo = honeycomb.NewHoneycombRepository(p, c, info)
			loggers = append(loggers, hcRepo)
		default:
			hcRepo = honeycomb.NewHoneycombRepository(p, c, info)
			loggers = append(loggers, hcRepo)
		}
	}

	loggerR := loggerRepo{
		p:       p,
		c:       c,
		span:    hcRepo.Span,
		loggers: loggers,
	}

	for _, opt := range options {
		opt(&loggerR)
	}

	return loggerR, func() {
		log.Print("span ended ", info.Path+info.FunctionName)
		hcRepo.Span.End()
	}
}

func (l *loggerRepo) Debug(msg string, fields map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Debug(msg, fields)
	}
}

func (l *loggerRepo) Info(msg string, fields map[string]interface{}) {

	log.Print("!!! info called")

	for _, logger := range l.loggers {
		log.Print("logger in Info")
		logger.Info(msg, fields)
	}
}

func (l *loggerRepo) Warn(msg string, fields map[string]interface{}) {

	for _, logger := range l.loggers {
		logger.Warn(msg, fields)
	}
}

func (l *loggerRepo) Error(msg string, fields map[string]interface{}) {

	for _, logger := range l.loggers {
		logger.Error(msg, fields)
	}
}

func (l *loggerRepo) Fatal(msg string, fields map[string]interface{}) {

	for _, logger := range l.loggers {
		logger.Fatal(msg, fields)
	}
}

func SetNewOtelContext() Option {
	return func(c *loggerRepo) {
		for _, logger := range c.loggers {
			if honeycombRepo, ok := logger.(*honeycomb.HoneycombRepo); ok {
				honeycombRepo.SetNewOtelContext()
			}
		}
	}
}

func (l *loggerRepo) SetContextFromSpan() {
	newCtx := trace.ContextWithSpan(l.c.Request.Context(), l.span)
	l.c.Set("otel-context", newCtx)
}

func (l *loggerRepo) End() {

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
