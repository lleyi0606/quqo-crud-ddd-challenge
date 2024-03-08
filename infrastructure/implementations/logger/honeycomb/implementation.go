package honeycomb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	loggerentity "products-crud/domain/entity/logger_entity"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// type HoneycombRepo struct {
// 	p *base.Persistence
// 	c *gin.Context
// 	// s *trace.Span
// }

type HoneycombRepo struct {
	p            *base.Persistence
	c            *gin.Context
	Span         trace.Span
	otel_context context.Context
	info         loggerentity.FunctionInfo
}

func NewHoneycombRepository(p *base.Persistence, c *gin.Context, info loggerentity.FunctionInfo) *HoneycombRepo {

	// Start a new span
	// tracer := otel.Tracer("") // honeycomb.io
	tracer := otel.GetTracerProvider().Tracer("")

	// Retrieve existing context or use the request context
	ctx, ctxFound := c.Get("otel-context")
	if !ctxFound {
		log.Println("otel-context NOT FOUND")
		ctx = c.Request.Context()
		c.Set("otel-context", ctx)
	} else {
		log.Println("otel-context FOUND")
	}

	// Start a new span with attributes
	context, span := tracer.Start(ctx.(context.Context), info.Path+info.FunctionName,
		trace.WithAttributes(
			attribute.String("Description", info.Description),
		))

	return &HoneycombRepo{p, c, span, context, info}
}

// func NewHoneycombRepository(p *base.Persistence, c *gin.Context, info loggerentity.FunctionInfo) *HoneycombRepo {
// 	_, callerInfo, _, _ := runtime.Caller(3)
// 	log.Printf("!!! new honeycomb called %s // %s", info.Path+info.FunctionName, callerInfo)

// 	// Start a new span
// 	tracer := otel.Tracer("") // honeycomb.io

// 	// Retrieve existing context or use the request context
// 	ctx, ctxFound := c.Get("otel-context")
// 	if !ctxFound {
// 		log.Println("otel-context NOT FOUND")
// 		ctx = c.Request.Context()
// 		c.Set("otel-context", ctx)
// 	} else {
// 		log.Println("otel-context FOUND")
// 	}

// 	// Check if the callerInfo is different
// 	storedCaller, callerInfoExists := c.Get("callerInfo")
// 	if callerInfoExists && storedCaller.(string) == callerInfo {
// 		log.Println("Reuse otel-context", storedCaller, "...", callerInfo)

// 		// Retrieve the previous context from the stored value
// 		if prevCtx, prevCtxExists := c.Get("previous-context"); prevCtxExists {
// 			// Start a new span using the previous context
// 			_, span := tracer.Start(prevCtx.(context.Context), info.Path+info.FunctionName,
// 				trace.WithAttributes(
// 					attribute.String("Description", info.Description),
// 				))
// 			c.Set("callerInfo", callerInfo)

// 			return &HoneycombRepo{p, c, span}
// 		}
// 	} else {
// 		log.Println("Store new otel-context", storedCaller, "...", callerInfo)

// 		// Start a new span using the current context
// 		context, span := tracer.Start(ctx.(context.Context), info.Path+info.FunctionName,
// 			trace.WithAttributes(
// 				attribute.String("Description", info.Description),
// 			))
// 		c.Set("previous-context", ctx)
// 		c.Set("otel-context", context)
// 		c.Set("callerInfo", callerInfo)

// 		return &HoneycombRepo{p, c, span}
// 	}

// 	return &HoneycombRepo{p, c, nil}
// }

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
	log.Print("otel context set from options")
	l.c.Set("otel-context", l.otel_context)
}

func (l *HoneycombRepo) End() {
	// Set otel-context immediately
	l.c.Set("otel-context", l.otel_context)
	log.Print("otel context set in End ", l.info.Path+l.info.FunctionName)

	// Defer a closure that calls l.Span.End()
	defer func() {
		log.Print("span ended in End() ", l.info.Path+l.info.FunctionName)
		l.Span.End()
	}()
}
