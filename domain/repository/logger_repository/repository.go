package logger_repository

import (
	"context"
	loggerentity "products-crud/domain/entity/span_entity"

	"go.opentelemetry.io/otel/trace"
)

type LoggerRepository interface {
	NewSpan(*loggerentity.Span) (context.Context, trace.Span)
	EndSpan(trace.Span)
	LogError(trace.Span, error)
}
