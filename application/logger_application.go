package application

import (
	"context"
	loggerentity "products-crud/domain/entity/span_entity"
	"products-crud/domain/repository/logger_repository"
	"products-crud/infrastructure/implementations/logger"
	base "products-crud/infrastructure/persistences"

	"go.opentelemetry.io/otel/trace"
)

type LoggerApp struct {
	p *base.Persistence
	c *context.Context
}

func NewLoggerApplication(p *base.Persistence, c *context.Context) logger_repository.LoggerRepository {
	return &LoggerApp{p, c}
}

func (l LoggerApp) NewSpan(span *loggerentity.Span) (context.Context, trace.Span) {
	loggerInventory := logger.NewLoggerRepository(l.p, l.c, "Honeycomb")
	return loggerInventory.NewSpan(span)
}

func (l LoggerApp) EndSpan(span trace.Span) {
	loggerInventory := logger.NewLoggerRepository(l.p, l.c, "Honeycomb")
	loggerInventory.EndSpan(span)
}

func (l LoggerApp) LogError(span trace.Span, err error) {
	loggerInventory := logger.NewLoggerRepository(l.p, l.c, "Honeycomb")
	loggerInventory.LogError(span, err)
}
