package honeycomb

import (
	"context"
	"log"
	loggerentity "products-crud/domain/entity/span_entity"
	"products-crud/domain/repository/logger_repository"
	base "products-crud/infrastructure/persistences"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type honeycombRepo struct {
	p *base.Persistence
	c *context.Context
	// s *trace.Span
}

func NewHoneycombRepository(p *base.Persistence, c *context.Context) logger_repository.LoggerRepository {
	return &honeycombRepo{p, c}
}

func (r honeycombRepo) NewSpan(span *loggerentity.Span) (context.Context, trace.Span) {
	log.Print("honeycomb NewSpan called")
	tracer := otel.Tracer("quqo")

	context, s := tracer.Start(*r.c, span.Path+span.FunctionName,
		trace.WithAttributes(
			attribute.String("Description", span.Description),
		))

	return context, s

}

func (r honeycombRepo) EndSpan(span trace.Span) {
	span.End()
}

func (r honeycombRepo) LogError(span trace.Span, err error) {
	span.RecordError(err)
}
