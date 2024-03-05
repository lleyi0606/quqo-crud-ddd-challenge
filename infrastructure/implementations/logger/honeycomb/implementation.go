package honeycomb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	loggerentity "products-crud/domain/entity/logger_entity"
	base "products-crud/infrastructure/persistences"
	"runtime"

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
	p    *base.Persistence
	c    *gin.Context
	Span trace.Span
	// Otel_context *context.Context
}

func NewHoneycombRepository(p *base.Persistence, c *gin.Context, info loggerentity.FunctionInfo) *HoneycombRepo {

	_, callerInfo, _, _ := runtime.Caller(3)
	log.Printf("!!! new honeycomb called %s // %s", info.Path+info.FunctionName, callerInfo)

	// Start a new span
	tracer := otel.Tracer("") // honeycomb.io

	var ctx context.Context
	if storedCtx, exists := c.Get("otel-context"); exists {
		log.Println("otel-context FOUND")
		ctx = storedCtx.(context.Context)
	} else {
		log.Println("otel-context NOT FOUND")
		ctx = c.Request.Context()
		c.Set("otel-context", ctx)
	}

	// context, span := p.Logger.HoneycombTracer.Start(*c, info.Path+info.FunctionName,
	context, span := tracer.Start(ctx, info.Path+info.FunctionName,
		trace.WithAttributes(
			attribute.String("Description", info.Description),
		))

	if _, exists := c.Get("otel-context"); exists {
		// Check if the callerInfo is different
		if storedCaller, ok := c.Get("callerInfo"); ok && storedCaller.(string) == callerInfo {
			log.Println("Reuse otel-context", storedCaller, "...", callerInfo)
		} else {
			log.Println("Store new otel-context", storedCaller, "...", callerInfo)
			c.Set("otel-context", context)
		}
	}

	// Store the current callerInfo for comparison
	c.Set("callerInfo", callerInfo)
	// c.Set("otel-context", context)

	// var ctx context.Context
	// if storedCtx, exists := c.Get("otel-context"); exists {
	// 	log.Println("otel-context FOUND")
	// 	ctx = storedCtx.(context.Context)
	// } else {
	// 	log.Println("otel-context NOT FOUND")
	// 	ctx = c.Request.Context()
	// }

	// // context, span := p.Logger.HoneycombTracer.Start(*c, info.Path+info.FunctionName,
	// context, span := tracer.Start(ctx, info.Path+info.FunctionName,
	// 	trace.WithAttributes(
	// 		attribute.String("Description", info.Description),
	// 	))

	// c.Set("otel-context", context)

	// log.Println("otel-context SET")

	// Defer the end of the span
	// defer span.End()

	// Return a new repository instance with the tracer and context
	return &HoneycombRepo{p, c, span}
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
