package base

import (
	"log"
	inventory_entity "products-crud/domain/entity/inventory_entity"
	"products-crud/domain/entity/opensearch_entity"
	product_entity "products-crud/domain/entity/product_entity"
	"products-crud/infrastructure/persistences/db"

	"go.uber.org/zap"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Persistence struct {
	ProductDb          *gorm.DB
	ProductRedisDb     *redis.Client
	ProductAlgoliaDb   *search.Index
	InventoryAlgoliaDb *search.Index
	SearchOpenSearchDb *opensearch_entity.OpenSearch
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
	algoliaPdtIndex, algoliaInventoryIndex, errAlgoliaProductR := db.NewProductAlgoliaDB()
	if errAlgoliaProductR != nil {
		zap.S().Error("ALGOLIA NOT INITIALIZED", "error", errAlgoliaProductR)
	}

	// OpenSearch
	opensearchIndex, errOpensearchR := db.NewProductOpenSearchDB()
	if errOpensearchR != nil {
		zap.S().Error("OPENSEARCH NOT INITIALIZED", "error", errOpensearchR)
	}

	return &Persistence{
		ProductDb:          productEngine.DB,
		ProductRedisDb:     redisClient,
		ProductAlgoliaDb:   algoliaPdtIndex,
		InventoryAlgoliaDb: algoliaInventoryIndex,
		SearchOpenSearchDb: opensearchIndex,
	}, nil
}

// closes the  database connection
func (p *Persistence) Close() error {
	pdtDB, errP := p.ProductDb.DB()
	if errP != nil {
		return errP
	}
	errDbClose := pdtDB.Close()
	if errDbClose != nil {
		return errDbClose
	}

	errRedisClose := p.ProductRedisDb.Close()
	if errRedisClose != nil {
		return errRedisClose
	}

	return nil
}

// This migrate all tables
func (p *Persistence) Automigrate() error {
	p.ProductDb.AutoMigrate(&inventory_entity.Inventory{})

	return p.ProductDb.AutoMigrate(&product_entity.Product{})
}
