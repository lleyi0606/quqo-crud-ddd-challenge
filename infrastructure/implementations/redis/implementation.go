package redis

import (
	"encoding/json"
	"errors"
	product_entity "products-crud/domain/entity/product_entity"
	redis_repository "products-crud/domain/repository/redis_respository"
	base "products-crud/infrastructure/persistences"
	"strings"
	"time"

	// "github.com/RediSearch/redisearch-go/redisearch"

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

func (r redisRepo) SearchName(key string, src interface{}) error {

	if r.p.ProductRedisDb == nil {
		return errors.New("REDIS NOT FOUND")
	}

	// get all the keys in Redis
	keys, err := r.p.ProductRedisDb.Keys("*").Result()
	if err != nil {
		zap.S().Error("1. Redis SearchName ERROR", "error", err, "key", key)
		return err
	}

	var products []product_entity.Product

	// Iterate through keys and check corresponding values for the keyword
	for _, key := range keys {
		val, err := r.p.ProductRedisDb.Get(key).Result()
		if err != nil {
			zap.S().Error("2. Redis SearchName ERROR", "error", err, "key", key)
		}

		var pdt product_entity.Product
		err = json.Unmarshal([]byte(val), &pdt)
		if err != nil {
			zap.S().Error("3. Redis SearchName ERROR", "error", err, "key", key)
		}

		if strings.Contains(pdt.Name, key) {
			products = append(products, pdt)
		}
	}

	src = products

	return nil

}

func NewRedisRepository(p *base.Persistence) redis_repository.RedisRepository {
	return &redisRepo{p}
}
