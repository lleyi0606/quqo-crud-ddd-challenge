package zap

import (
	"log"

	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type zapRepo struct {
	p *base.Persistence
	c *gin.Context
}

func NewZapRepository(p *base.Persistence, c *gin.Context) *zapRepo {
	log.Print("!!! new zap called")

	return &zapRepo{p: p, c: c}
}

func (l *zapRepo) Debug(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.p.Logger.ZapLogger.Debug("", zap.Any("args", fields))
}

func (l *zapRepo) Info(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)
	log.Print("!!! zap info called")

	l.p.Logger.ZapLogger.Info(msg, zap.Any("args", fields))
}

func (l *zapRepo) Warn(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.p.Logger.ZapLogger.Warn("", zap.Any("args", fields))
}

func (l *zapRepo) Error(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.p.Logger.ZapLogger.Error("", zap.Any("args", fields))
}

func (l *zapRepo) Fatal(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.p.Logger.ZapLogger.Fatal("", zap.Any("args", fields))
}

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
