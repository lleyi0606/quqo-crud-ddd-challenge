package base

import (
	"log"
	category_entity "products-crud/domain/entity/category_entity"
	customer_entity "products-crud/domain/entity/customer_entity"
	image_entity "products-crud/domain/entity/image_entity"
	inventory_entity "products-crud/domain/entity/inventory_entity"
	order_entity "products-crud/domain/entity/order_entity"
	orderedItem_entity "products-crud/domain/entity/orderedItem_entity"
	product_entity "products-crud/domain/entity/product_entity"

	"products-crud/infrastructure/persistences/db"

	"go.uber.org/zap"

	storage_go "github.com/supabase-community/storage-go"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
	"gorm.io/gorm"
)

type Persistence struct {
	ProductDb          *gorm.DB
	ProductRedisDb     *redis.Client
	ProductAlgoliaDb   *search.Index
	InventoryAlgoliaDb *search.Index
	SearchOpenSearchDb *opensearch.Client
	ImageSupabaseDB    *storage_go.Client
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

	// Supabase
	supabaseEngine, errSupabase := db.NewImageSupabaseDB()
	if errSupabase != nil {
		zap.S().Error("SUPABASES NOT INITIALIZED", "error", errSupabase)
	}

	return &Persistence{
		ProductDb:          productEngine.DB,
		ProductRedisDb:     redisClient,
		ProductAlgoliaDb:   algoliaPdtIndex,
		InventoryAlgoliaDb: algoliaInventoryIndex,
		SearchOpenSearchDb: opensearchIndex,
		ImageSupabaseDB:    supabaseEngine.Client,
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

// This migrates all tables
func (p *Persistence) Automigrate() error {
	p.ProductDb.AutoMigrate(&inventory_entity.Inventory{})
	p.ProductDb.AutoMigrate(&image_entity.Image{})
	p.ProductDb.AutoMigrate(&category_entity.Category{})
	p.ProductDb.AutoMigrate(&customer_entity.Customer{})
	p.ProductDb.AutoMigrate(&order_entity.Order{})
	p.ProductDb.AutoMigrate(&orderedItem_entity.OrderedItem{})
	return p.ProductDb.AutoMigrate(&product_entity.Product{})
}
