package db

import (
	loggerentity "products-crud/domain/entity/logger_entity"

	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

// func newConsoleExporter() (sdktrace.SpanExporter, error) {
// 	// Create a new console exporter
// 	sdktrace.New
// 	exporter, err := stdout.NewExporter(stdout.WithPrettyPrint())
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create console exporter: %w", err)
// 	}
// 	return exporter, nil
// }

func newTraceProvider() *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	// r, err := resource.Merge(
	// 	resource.Default(),
	// 	resource.NewWithAttributes(
	// 		semconv.SchemaURL,
	// 		semconv.ServiceName("ExampleService"),
	// 	),
	// )

	// if err != nil {
	// 	panic(err)
	// }

	return sdktrace.NewTracerProvider(
	// sdktrace.WithBatcher(exp),
	// sdktrace.WithResource(r),
	)
}

func NewLoggerDB() (*loggerentity.Logger, error) {

	// honeycomb for tracing
	// tracer := otel.Tracer("backend") // honeycomb.io
	tp := newTraceProvider()

	// Handle shutdown properly so nothing leaks.
	// defer func() { _ = tp.Shutdown(context.Background()) }()

	otel.SetTracerProvider(tp)
	// if tracePovider == nil {
	// 	zap.S().Errorw("Tracer provider initialise error")
	// 	return nil, errors.New("error tracer connection")
	// }

	// zap for logging
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	undo := zap.ReplaceGlobals(zapLogger)
	defer undo()

	logger := loggerentity.Logger{
		HoneycombTracer: tp,
		ZapLogger:       zapLogger,
	}

	return &logger, nil

}
