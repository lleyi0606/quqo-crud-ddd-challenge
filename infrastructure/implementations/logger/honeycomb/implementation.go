package honeycomb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type HoneycombRepo struct {
	c            *gin.Context
	Span         trace.Span
	otel_context context.Context
	// info         loggerentity.FunctionInfo
}

func NewHoneycombRepository() *HoneycombRepo {

	return &HoneycombRepo{nil, nil, nil}
}

func (h *HoneycombRepo) Start(c *gin.Context, functionPath string, fields map[string]interface{}) trace.Span {
	// Start a new span
	tracer := otel.Tracer("") // honeycomb.io

	// Retrieve existing context or use the request context
	ctx, ctxFound := c.Get("otel-context")
	if !ctxFound {
		ctx = c.Request.Context()
		c.Set("otel-context", ctx)
	}

	// Start a new span with attributes
	commonAttributes := getCommonAttributes(c)
	attributes := append([]attribute.KeyValue{attribute.String("FunctionPath", functionPath)}, commonAttributes...)
	context, span := tracer.Start(ctx.(context.Context), functionPath,
		trace.WithAttributes(
			attributes...,
		))

	h.c = c
	h.Span = span
	h.otel_context = context

	return span
}

func (l *HoneycombRepo) Debug(msg string, fields map[string]interface{}) {
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Debug:", err)
	}

	commonAttributes := getCommonAttributes(l.c)
	attributes := append([]attribute.KeyValue{attribute.String("level", "debug"),
		attribute.String("data", string(jsonObj))}, commonAttributes...)
	l.Span.AddEvent(msg,
		trace.WithAttributes(
			attributes...,
		))
}

func (l *HoneycombRepo) Info(msg string, fields map[string]interface{}) {
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Info:", err)
	}
	commonAttributes := getCommonAttributes(l.c)
	attributes := append([]attribute.KeyValue{attribute.String("level", "info"),
		attribute.String("data", string(jsonObj))}, commonAttributes...)
	l.Span.AddEvent(msg,
		trace.WithAttributes(
			attributes...,
		))
}

func (l *HoneycombRepo) Warn(msg string, fields map[string]interface{}) {
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Warn:", err)
	}
	commonAttributes := getCommonAttributes(l.c)
	attributes := append([]attribute.KeyValue{attribute.String("level", "warn"),
		attribute.String("data", string(jsonObj))}, commonAttributes...)
	l.Span.AddEvent(msg,
		trace.WithAttributes(
			attributes...,
		))
}

func (l *HoneycombRepo) Error(msg string, fields map[string]interface{}) {
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Error:", err)
		return
	}
	commonAttributes := getCommonAttributes(l.c)
	attributes := append([]attribute.KeyValue{attribute.String("level", "error"),
		attribute.String("data", string(jsonObj))}, commonAttributes...)
	l.Span.AddEvent(msg,
		trace.WithAttributes(
			attributes...,
		))
	l.Span.RecordError(fmt.Errorf("Error: %s", string(jsonObj)))
	l.Span.SetStatus(codes.Error, msg)
}

func (l *HoneycombRepo) Fatal(msg string, fields map[string]interface{}) {
	// Convert the fields to JSON
	jsonObj, err := json.Marshal(fields)
	if err != nil {
		log.Println("Error marshaling data to JSON in Fatal:", err)
	}

	// Add an event with the error details
	commonAttributes := getCommonAttributes(l.c)
	attributes := append([]attribute.KeyValue{attribute.String("level", "fatal"),
		attribute.String("data", string(jsonObj))}, commonAttributes...)
	l.Span.AddEvent(msg,
		trace.WithAttributes(
			attributes...,
		))

	// Record an error to make sure it's captured by OpenTelemetry
	l.Span.RecordError(fmt.Errorf("Fatal error: %s", msg))
	l.Span.SetStatus(codes.Error, msg)

	// Terminate the application
	os.Exit(1)
}

func (l *HoneycombRepo) SetNewOtelContext() {
	l.c.Set("otel-context", l.otel_context)
}

func (l *HoneycombRepo) UseGivenSpan(span trace.Span) {
	l.Span = span
}

func getCommonAttributes(c *gin.Context) []attribute.KeyValue {
	// Get caller information (file name and line number)
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "unknown"
		line = 0
	} else {
		// Extract only the file name from the full path
		fileParts := strings.Split(file, "/")
		file = fileParts[len(fileParts)-1]
	}

	attributes := []attribute.KeyValue{
		attribute.String("IP Address", c.ClientIP()),
		attribute.String("Environment", os.Getenv("ENV")),
		attribute.String("CallerFile", file),
		attribute.Int("CallerLine", line),
	}

	return attributes
}

// func (l *HoneycombRepo) End() {
// 	l.Span.End()
// }

// func getSpanID(ctx context.Context) trace.SpanID {
// 	// Retrieve the span from the context
// 	span := trace.SpanFromContext(ctx)

// 	// Access the span ID from the span
// 	spanContext := span.SpanContext()
// 	spanID := spanContext.SpanID()

// 	return spanID
// }
