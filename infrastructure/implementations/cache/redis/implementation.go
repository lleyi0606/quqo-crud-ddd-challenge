package redis

import (
	"encoding/json"
	"errors"
	cache_repository "products-crud/domain/repository/cache_repository"
	base "products-crud/infrastructure/persistences"
	"time"

	"go.uber.org/zap"
)

type redisRepo struct {
	p *base.Persistence
}

func (r redisRepo) SetKey(key string, value interface{}, expiration time.Duration) error {

	if r.p.ProductRedisDb == nil {
		return errors.New("REDIS NOT FOUND")
	}

	cacheEntry, err := json.Marshal(value)
	if err != nil {
		zap.S().Errorw("1. Redis SetKey ERROR", "error", err, "key", key, "value", value)
		return err
	}
	err = r.p.ProductRedisDb.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		zap.S().Errorw("2. Redis SetKey ERROR", "error", err, "key", key, "value", value)
		return err
	}
	return nil

}

func (r redisRepo) GetKey(key string, src interface{}) error {

	if r.p.ProductRedisDb == nil {
		return errors.New("REDIS NOT FOUND")
	}

	val, err := r.p.ProductRedisDb.Get(key).Result()
	if err != nil {
		zap.S().Error("1. Redis GetKey ERROR", "error", err, "key", key)
		return err
	}

	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		zap.S().Error("2. Redis GetKey ERROR", "error", err, "key", key)
		return err
	}

	return nil

}

// func (r redisRepo) ProductExists(id uint64) (int64, error) {
// 	exists, err := r.p.ProductRedisDb.Exists(fmt.Sprintf("%s%d", redis_entity.RedisProductData, id)).Result()
// 	if err != nil {
// 		zap.S().Error("Redis ProductExists ERROR", "error", err, "key", id)
// 		return 0, err
// 	}
// 	return exists, nil
// }

func (r redisRepo) DeleteRecord(key string) error {
	err := r.p.ProductRedisDb.Del(key).Err()
	if err != nil {
		zap.S().Error("Redis DeleteRecord ERROR", "error", err, "key", key)
		return err
	}
	return nil
}

func NewRedisRepository(p *base.Persistence) cache_repository.CacheRepository {
	return &redisRepo{p}
}
