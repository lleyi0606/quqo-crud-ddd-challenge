package zap

import (
	"log"

	"github.com/gin-gonic/gin"
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

func (z *zapRepo) Start(c *gin.Context, functionPath string, fields map[string]interface{}) func() {
	return func() {}
}

func (l *zapRepo) Debug(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.zapLogger.Debug("", zap.Any("args", fields))
}

func (l *zapRepo) Info(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)
	log.Print("!!! zap info called")

	l.zapLogger.Info(msg, zap.Any("args", fields))
}

func (l *zapRepo) Warn(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.zapLogger.Warn("", zap.Any("args", fields))
}

func (l *zapRepo) Error(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.zapLogger.Error("", zap.Any("args", fields))
}

func (l *zapRepo) Fatal(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.zapLogger.Fatal("", zap.Any("args", fields))
}

func (l *zapRepo) End() {}

// func (l *zapRepo) addContextCommonFields(fields map[string]interface{}) {
// 	if fields == nil || l.ctx == nil {
// 		return // Do nothing if either fields or l.ctx is nil
// 	}

// 	if l.ctx != nil {
// 		for k, v := range l.ctx.Value("commonFields").(map[string]interface{}) {
// 			if _, ok := fields[k]; !ok {
// 				fields[k] = v
// 			}
// 		}
// 	}
// }
