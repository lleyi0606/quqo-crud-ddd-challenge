package db

import (
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func NewHoneycombDB() (*trace.Tracer, error) {

	tracer := otel.Tracer("quqo")

	if tracer == nil {
		zap.S().Errorw("Tracer initialise error")
		return nil, errors.New("error redis connection")
	}

	return &tracer, nil

}
