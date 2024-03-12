package zap

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type zapRepo struct {
	c         *gin.Context
	zapLogger *zap.Logger
}

func NewZapRepository() *zapRepo {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	undo := zap.ReplaceGlobals(zapLogger)
	defer undo()

	return &zapRepo{nil, zapLogger}
}

func (z *zapRepo) Start(c *gin.Context, functionPath string, fields map[string]interface{}) trace.Span {
	return nil
}

func (l *zapRepo) Debug(msg string, fields map[string]interface{}) {

	l.zapLogger.Debug("", zap.Any("args", fields))
}

func (l *zapRepo) Info(msg string, fields map[string]interface{}) {

	l.zapLogger.Info(msg, zap.Any("args", fields))
}

func (l *zapRepo) Warn(msg string, fields map[string]interface{}) {

	l.zapLogger.Warn("", zap.Any("args", fields))
}

func (l *zapRepo) Error(msg string, fields map[string]interface{}) {

	l.zapLogger.Error("", zap.Any("args", fields))
}

func (l *zapRepo) Fatal(msg string, fields map[string]interface{}) {

	l.zapLogger.Fatal("", zap.Any("args", fields))
}

// func (l *zapRepo) End() {}
