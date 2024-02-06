package cache

import (
	cache_repository "products-crud/domain/repository/cache_repository"
	"products-crud/infrastructure/implementations/cache/redis"
	base "products-crud/infrastructure/persistences"
)

const (
	Redis = "redis"
)

func NewCacheRepository(p *base.Persistence, provider string) cache_repository.CacheRepository {
	switch provider {
	case Redis:
		return redis.NewRedisRepository(p)
	default:
		return redis.NewRedisRepository(p)
	}
}
