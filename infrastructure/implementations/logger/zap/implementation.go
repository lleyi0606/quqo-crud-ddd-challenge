package zap

import (
	"context"
	"log"

	base "products-crud/infrastructure/persistences"

	"go.uber.org/zap"
)

type zapRepo struct {
	logger *zap.Logger
	ctx    context.Context
}

func NewZapRepository(p *base.Persistence, c *context.Context) *zapRepo {
	log.Print("!!! new zap called")

	return &zapRepo{logger: p.Logger.ZapLogger, ctx: *c}
}

func (l *zapRepo) Debug(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.logger.Debug("", zap.Any("args", fields))
}

func (l *zapRepo) Info(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)
	log.Print("!!! zap info called")

	l.logger.Info(msg, zap.Any("args", fields))
}

func (l *zapRepo) Warn(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.logger.Warn("", zap.Any("args", fields))
}

func (l *zapRepo) Error(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.logger.Error("", zap.Any("args", fields))
}

func (l *zapRepo) Fatal(msg string, fields map[string]interface{}) {
	// l.addContextCommonFields(fields)

	l.logger.Fatal("", zap.Any("args", fields))
}

func (l *zapRepo) addContextCommonFields(fields map[string]interface{}) {
	if fields == nil || l.ctx == nil {
		return // Do nothing if either fields or l.ctx is nil
	}

	if l.ctx != nil {
		for k, v := range l.ctx.Value("commonFields").(map[string]interface{}) {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}
	}
}
