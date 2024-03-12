package honeycomb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type HoneycombRepo struct {
	c            *gin.Context
	Span         trace.Span
	Otel_context context.Context
	// info         loggerentity.FunctionInfo
}

func NewHoneycombRepository() *HoneycombRepo {

	return &HoneycombRepo{nil, nil, nil}
}

func (h *HoneycombRepo) Start(c *gin.Context, functionPath string, fields map[string]interface{}) trace.Span {
	log.Print("started in ", functionPath)
	// Start a new span
	tracer := otel.Tracer("") // honeycomb.io
	// tracer := otel.GetTracerProvider().Tracer("")

	// Retrieve existing context or use the request context
	ctx, ctxFound := c.Get("otel-context")
	if !ctxFound {
		log.Println("otel-context NOT FOUND")
		ctx = c.Request.Context()
		c.Set("otel-context", ctx)
	} else {
		log.Println("otel-context FOUND", getSpanID(ctx.(context.Context)))
	}

	// Start a new span with attributes
	context, span := tracer.Start(ctx.(context.Context), functionPath,
		trace.WithAttributes(
			attribute.String("FunctionPath", functionPath),
		))

	h.c = c
	h.Span = span
	h.Otel_context = context

	// return func() {
	// 	span.End()
	// }
	return span
}

func (l *HoneycombRepo) Debug(msg string, fields map[string]interface{}) {

	log.Print("!!! debug in honeycomb")
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Debug:", err)
	}
	l.Span.AddEvent(msg, trace.WithAttributes(
		attribute.String("level", "debug"),
		attribute.String("data", string(jsonObj)),
	))
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

func (l *HoneycombRepo) SetNewOtelContext() {
	// log.Print("otel context set from option ", getSpanID(l.Otel_context))
	l.c.Set("otel-context", l.Otel_context)
}

func (l *HoneycombRepo) End() {
	// Set otel-context immediately
	// l.c.Set("otel-context", l.otel_context)
	// log.Print("otel context set in End ", l.info.Path+l.info.FunctionName)

	// Defer a closure that calls l.Span.End()
	// defer func() {
	// log.Print("span ended in End() ", l.info.Path+l.info.FunctionName)
	l.Span.End()
	// }()
}

func getSpanID(ctx context.Context) trace.SpanID {
	// Retrieve the span from the context
	span := trace.SpanFromContext(ctx)

	// Access the span ID from the span
	spanContext := span.SpanContext()
	spanID := spanContext.SpanID()

	return spanID
}
