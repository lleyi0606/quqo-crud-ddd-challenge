package honeycomb

import (
	"context"
	"encoding/json"
	"log"
	loggerentity "products-crud/domain/entity/logger_entity"
	base "products-crud/infrastructure/persistences"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// type HoneycombRepo struct {
// 	p *base.Persistence
// 	c *context.Context
// 	// s *trace.Span
// }

type HoneycombRepo struct {
	logger *trace.Tracer
	Ctx    context.Context
	Span   trace.Span
}

func NewHoneycombRepository(p *base.Persistence, c *context.Context, info loggerentity.FunctionInfo) *HoneycombRepo {

	log.Printf("!!! new honeycomb called %s", info.Path+info.FunctionName)

	// Start a new span
	tracer := otel.Tracer("") // honeycomb.io

	// context, span := p.Logger.HoneycombTracer.Start(*c, info.Path+info.FunctionName,
	context, span := tracer.Start(*c, info.Path+info.FunctionName,
		trace.WithAttributes(
			attribute.String("Description", info.Description),
		))

	// Defer the end of the span
	// defer span.End()

	// Return a new repository instance with the tracer and context
	return &HoneycombRepo{p.Logger.HoneycombTracer, context, span}

}

func (l *HoneycombRepo) Debug(msg string, fields map[string]interface{}) {

	// l.logger.Debug("", zap.Any("args", fields))
}

func (l *HoneycombRepo) Info(msg string, fields map[string]interface{}) {

	// l.logger.Info("", zap.Any("args", fields))

	// span.AddEvent("Sending JSON data to Honeycomb", trace.WithAttributes(
	// 	attribute.String("json_data", string(orderJSON)),
	// ))

	// Convert order to JSON

	log.Print("!!! info in honeycomb")
	orderJSON, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling order to JSON:", err)
	}
	l.Span.AddEvent("Sending JSON data to Honeycomb", trace.WithAttributes(
		attribute.String("json_data", string(orderJSON)),
	))
}

func (l *HoneycombRepo) Warn(msg string, fields map[string]interface{}) {

	// l.logger.Warn("", zap.Any("args", fields))
}

func (l *HoneycombRepo) Error(msg string, fields map[string]interface{}) {

	// l.logger.Error("", zap.Any("args", fields))
}

func (l *HoneycombRepo) Fatal(msg string, fields map[string]interface{}) {

	// l.logger.Fatal("", zap.Any("args", fields))
}

// func (r HoneycombRepo) NewSpan(span *loggerentity.Span) (context.Context, trace.Span) {
// 	log.Print("honeycomb NewSpan called")
// 	tracer := otel.Tracer("quqo")

// 	context, s := tracer.Start(*r.c, span.Path+span.FunctionName,
// 		trace.WithAttributes(
// 			attribute.String("Description", span.Description),
// 		))

// 	return context, s

// }

// func (r HoneycombRepo) EndSpan(span trace.Span) {
// 	span.End()
// }

// func (r HoneycombRepo) LogError(span trace.Span, err error) {
// 	span.RecordError(err)
// }
