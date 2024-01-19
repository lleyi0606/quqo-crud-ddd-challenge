package base

import (
	"log"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/infrastructure/persistences/db"

	"go.uber.org/zap"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Persistence struct {
	ProductDb        *gorm.DB
	ProductRedisDb   *redis.Client
	ProductAlgoliaDb *search.Index
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

	// Product Algolia engine
	algoliaIndex, errAlgoliaProductR := db.NewProductAlgoliaDB()
	if errAlgoliaProductR != nil {
		zap.S().Error("ALGOLIA NOT INITIALIZED", "error", errAlgoliaProductR)
	}

	return &Persistence{
		ProductDb:        productEngine.DB,
		ProductRedisDb:   redisClient,
		ProductAlgoliaDb: algoliaIndex,
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
