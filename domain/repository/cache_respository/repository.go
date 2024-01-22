package cache_repository

import "time"

type CacheRepository interface {
	SetKey(key string, value interface{}, expiration time.Duration) error
	GetKey(key string, src interface{}) error
	ProductExists(uint64) (int64, error)
	DeleteProduct(uint64) error
}
