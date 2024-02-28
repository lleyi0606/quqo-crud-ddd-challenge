package logger

import (
	"context"
	"products-crud/domain/repository/logger_repository"
	"products-crud/infrastructure/implementations/logger/honeycomb"
	base "products-crud/infrastructure/persistences"
)

const (
	Honeycomb = "redis"
)

func NewLoggerRepository(p *base.Persistence, c *context.Context, provider string) logger_repository.LoggerRepository {
	switch provider {
	case Honeycomb:
		return honeycomb.NewHoneycombRepository(p, c)
	default:
		return honeycomb.NewHoneycombRepository(p, c)
	}
}
