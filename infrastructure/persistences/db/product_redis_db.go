package db

import (
	"errors"
	"products-crud/infrastructure/config"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func NewProductRedisDB() (*redis.Client, error) {

	// connection DB
	// DbHost := os.Getenv("DB_HOST_REDIS")
	// DbPassword := os.Getenv("DB_PASSWORD_REDIS")

	DbHost := config.Configuration.GetString("redis.dev.host")
	DbPassword := config.Configuration.GetString("redis.dev.pass")

	c := redis.NewClient(&redis.Options{
		Addr:     DbHost,
		Password: DbPassword, // no password set
		DB:       0,          // use default DB
	})

	if err := c.Ping().Err(); err != nil {
		zap.S().Errorw("Redis Initialize err", "addr", DbHost, "err", err)
		return nil, errors.New("error redis connection" + err.Error())
	}

	return c, nil

}
