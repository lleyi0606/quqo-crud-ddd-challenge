package redis_repository

import "time"

type RedisRepository interface {
	SetKey(key string, value interface{}, expiration time.Duration) error
	GetKey(key string, src interface{}) error
	SearchName(key string, src interface{}) error
}
