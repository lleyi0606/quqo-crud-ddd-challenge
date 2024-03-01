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
	p       *base.Persistence
	c       *gin.Context
	loggers []logger_repository.LoggerRepository
}

// func NewLoggerRepositories(p *base.Persistence, c *gin.Context, info loggerentity.FunctionInfo, providers ...string) []logger_repository.LoggerRepository {
func NewLoggerRepositories(p *base.Persistence, c *gin.Context, info loggerentity.FunctionInfo, providers ...string) loggerRepo {

	log.Print("!!! new logger repo called")
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

	return loggerRepo{
		p:       p,
		c:       c,
		loggers: loggers,
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

func (l *loggerRepo) End() {
	var ctx context.Context
	if storedCtx, exists := l.c.Get("otel-context"); exists {
		log.Println("otel-context FOUND in End")
		ctx = storedCtx.(context.Context)
	} else {
		log.Println("otel-context NOT FOUND in End")
		ctx = l.c.Request.Context()
	}

	span := trace.SpanFromContext(ctx)
	defer span.End()
}
