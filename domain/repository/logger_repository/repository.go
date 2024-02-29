package logger_repository

type LoggerRepository interface {
	// NewSpan(*loggerentity.Span) (context.Context, trace.Span)
	// EndSpan(trace.Span)
	// LogError(trace.Span, error)

	Debug(msg string, fields map[string]interface{})
	Info(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Fatal(msg string, fields map[string]interface{})
}
