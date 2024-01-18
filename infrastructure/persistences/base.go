package base

import (
	"log"
	"products-crud/domain/entity"
	"products-crud/infrastructure/persistences/db"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Persistence struct {
	ProductDb      *gorm.DB
	ProductRedisDb *redis.Client
}

func NewPersistence() (*Persistence, error) {

	// Product engine
	productEngine, errProductE := db.NewProductDB()
	if errProductE != nil {
		log.Fatal(errProductE)
	}

	// Product Redis engine
	redisClient, errRedisProductR := db.NewProductRedisDB()
	if errRedisProductR != nil {
		zap.S().Error("REDIS NOT INITIALIZED", "error", errRedisProductR)
	}
	return &Persistence{
		ProductDb:      productEngine.DB,
		ProductRedisDb: redisClient,
	}, nil
}

// closes the  database connection
func (p *Persistence) Close() error {
	pdtDB, errQ := p.ProductDb.DB()
	if errQ != nil {
		return errQ
	}
	errDbClose := pdtDB.Close()
	if errDbClose != nil {
		return errDbClose
	}

	return nil
}

// This migrate all tables
func (p *Persistence) Automigrate() error {
	return p.ProductDb.AutoMigrate(&entity.Product{})
}
