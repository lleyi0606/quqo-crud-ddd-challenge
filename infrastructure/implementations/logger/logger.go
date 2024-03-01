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

	"go.opentelemetry.io/otel/trace"
)

const (
	Honeycomb = "HONEYCOMB"
	Zap       = "ZAP"
)

type orderRepo struct {
	p       *base.Persistence
	Context *context.Context
	loggers []logger_repository.LoggerRepository
}

// func NewLoggerRepositories(p *base.Persistence, c *context.Context, info loggerentity.FunctionInfo, providers ...string) []logger_repository.LoggerRepository {
func NewLoggerRepositories(p *base.Persistence, c *context.Context, info loggerentity.FunctionInfo, providers ...string) orderRepo {

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

	return orderRepo{
		p:       p,
		Context: &hcRepo.Ctx,
		loggers: loggers,
	}
}

func (l *orderRepo) Debug(msg string, fields map[string]interface{}) {
	for _, logger := range l.loggers {
		logger.Debug(msg, fields)
	}
}

func (l *orderRepo) Info(msg string, fields map[string]interface{}) {

	log.Print("!!! info called")

	for _, logger := range l.loggers {
		log.Print("logger in Info")
		logger.Info(msg, fields)
	}
}

func (l *orderRepo) Warn(msg string, fields map[string]interface{}) {

	for _, logger := range l.loggers {
		logger.Warn(msg, fields)
	}
}

func (l *orderRepo) Error(msg string, fields map[string]interface{}) {

	for _, logger := range l.loggers {
		logger.Error(msg, fields)
	}
}

func (l *orderRepo) Fatal(msg string, fields map[string]interface{}) {

	for _, logger := range l.loggers {
		logger.Fatal(msg, fields)
	}
}

func (l *orderRepo) End() {
	span := trace.SpanFromContext(*l.Context)
	defer span.End()
}
