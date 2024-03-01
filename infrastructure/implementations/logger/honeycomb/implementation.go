package honeycomb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	loggerentity "products-crud/domain/entity/logger_entity"
	base "products-crud/infrastructure/persistences"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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

	log.Print("!!! info in honeycomb")
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Info:", err)
	}
	l.Span.AddEvent(msg, trace.WithAttributes(
		attribute.String("level", "info"),
		attribute.String("data", string(jsonObj)),
	))
}

func (l *HoneycombRepo) Warn(msg string, fields map[string]interface{}) {

	log.Print("!!! warn in honeycomb")
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Warn:", err)
	}
	l.Span.AddEvent(msg, trace.WithAttributes(
		attribute.String("level", "warn"),
		attribute.String("data", string(jsonObj)),
	))
}

func (l *HoneycombRepo) Error(msg string, fields map[string]interface{}) {

	log.Print("!!! info in honeycomb")
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Error:", err)
		// l.Span.RecordError(fmt.Errorf("Error marshaling data to JSON: %v", err))
		return
	}
	l.Span.AddEvent(msg, trace.WithAttributes(
		attribute.String("level", "error"),
		attribute.String("data", string(jsonObj)),
	))
	l.Span.RecordError(fmt.Errorf("Error: %s", string(jsonObj)))
	l.Span.SetStatus(codes.Error, msg)
}

func (l *HoneycombRepo) Fatal(msg string, fields map[string]interface{}) {

	// l.logger.Fatal("", zap.Any("args", fields))
	log.Print("!!! fatal in honeycomb")
	// Convert the fields to JSON
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Fatal:", err)
	}

	// Add an event with the error details
	l.Span.AddEvent(msg, trace.WithAttributes(
		attribute.String("level", "fatal"),
		attribute.String("data", string(jsonObj)),
	))

	// Record an error to make sure it's captured by OpenTelemetry
	l.Span.RecordError(fmt.Errorf("Fatal error: %s", msg))
	l.Span.SetStatus(codes.Error, msg)

	// Terminate the application
	os.Exit(1)
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
