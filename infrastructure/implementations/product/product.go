package product

import (
	"errors"
	"fmt"
	"log"
	loggerentity "products-crud/domain/entity/logger_entity"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/entity/redis_entity"
	repository "products-crud/domain/repository/product_respository"
	"products-crud/infrastructure/implementations/cache"
	"products-crud/infrastructure/implementations/logger"
	"products-crud/infrastructure/implementations/search"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productRepo struct {
	p *base.Persistence
	c *gin.Context
}

func NewProductRepository(p *base.Persistence, c *gin.Context) repository.ProductRepository {
	return &productRepo{p, c}
}

// productRepo implements the repository.ProductRepository interface

func (r productRepo) AddProduct(pdt *entity.Product) (*entity.Product, error) {
	log.Println("Adding new product ", pdt.Name, "...")

	info := loggerentity.FunctionInfo{
		FunctionName: "AddProduct",
		Path:         "infrastructure/implementations/",
		Description:  "Adds product to SQL database",
		Body:         pdt,
	}
	logger := logger.NewLoggerRepositories(r.p, r.c, info, "honeycomb", "zap")
	defer logger.End()

	if err := r.p.ProductDb.Debug().Create(&pdt).Error; err != nil {
		// logger.Error(err.Error(), map[string]interface{}{"input": pdt})
		return nil, err
	}

	// add to search repo
	searchRepo := search.NewSearchRepository(r.p, "algolia")
	err := searchRepo.AddProduct(pdt)
	if err != nil {
		return nil, err
	}

	log.Println(pdt.Name, " created.")
	return pdt, nil
}

func (r productRepo) GetProduct(id uint64) (*entity.Product, error) {
	var pdt *entity.Product

	cacheRepo := cache.NewCacheRepository(r.p, "redis")
	_ = cacheRepo.GetKey(fmt.Sprintf("%s%d", redis_entity.RedisProductData, id), &pdt)

	if pdt == nil {
		err := r.p.ProductDb.Debug().Unscoped().Preload("Inventory").Where("product_id = ?", id).Take(&pdt).Error
		// err := r.p.ProductDb.Debug().Where("product_id = ?", id).Take(&pdt).Error
		if err != nil {
			return nil, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		_ = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisProductData, id), pdt, redis_entity.RedisExpirationGlobal)
	}

	return pdt, nil
}

func (r productRepo) GetProducts() ([]entity.Product, error) {
	var pdts []entity.Product
	err := r.p.ProductDb.Debug().Find(&pdts).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return pdts, nil
}

func (r productRepo) UpdateProduct(pdt *entity.Product) (*entity.Product, error) {
	// err := r.p.ProductDb.Debug().Model(&entity.Product{}).Where("id = ?", pdt.ID).Updates(pdt).Error
	err := r.p.ProductDb.Debug().Where("product_id = ?", pdt.ProductID).Updates(&pdt).Error

	if err != nil {
		return nil, err
	}

	// update cache
	cacheRepo := cache.NewCacheRepository(r.p, "redis")
	err = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisProductData, pdt.ProductID), &pdt, redis_entity.RedisExpirationGlobal)
	if err != nil {
		return nil, err
	}

	// update search repo
	// searchRepo := search.NewSearchRepository(r.p, "opensearch")
	// err = searchRepo.UpdateProduct(pdt)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return pdt, nil
}

func (r productRepo) DeleteProduct(id uint64) (*entity.Product, error) {
	var pdt entity.Product
	res := r.p.ProductDb.Debug().Where("product_id = ?", id).Delete(&pdt)
	if res.Error != nil {
		return nil, res.Error
	}
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return nil, errors.New("product not found")
	// }
	if res.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}

	// delete from inventory too
	// inventoryRepo := inventory.NewInventoryRepository(r.p)
	// _, err := inventoryRepo.DeleteInventory(id)
	// if err != nil {
	// 	return nil, err
	// }

	// search repo
	searchRepo := search.NewSearchRepository(r.p, "algolia")
	err := searchRepo.DeleteProduct(id)
	if err != nil {
		return nil, err
	}

	// update cache
	cacheRepo := cache.NewCacheRepository(r.p, "redis")
	err = cacheRepo.DeleteRecord(fmt.Sprintf("%s%d", redis_entity.RedisProductData, id))
	if err != nil {
		return nil, err
	}

	return &pdt, nil
}

func (r productRepo) SearchProducts(str string) ([]entity.Product, error) {

	// new search repo
	searchRepo := search.NewSearchRepository(r.p, "algolia")
	pdts, err := searchRepo.SearchProducts(str)
	if err != nil {
		return nil, err
	}
	return pdts, nil
}

func (r productRepo) CalculateProductPriceByQuantity(id uint64, qty int) (float64, float64, error) {
	pdt, err := r.GetProduct(id)
	if err != nil {
		return 0, 0, err
	}

	return pdt.Price, pdt.Price * float64(qty), nil
}

func (r productRepo) CalculateProductPriceByQuantityTx(tx *gorm.DB, id uint64, qty int) (float64, float64, error) {

	info := loggerentity.FunctionInfo{
		FunctionName: "CalculateProductPriceByQuantityTx",
		Path:         "infrastructure/implementations/",
		Description:  "Calculate price of a product",
		Body:         nil,
	}
	logger := logger.NewLoggerRepositories(r.p, r.c, info, "honeycomb", "zap")
	defer logger.End()

	pdt, err := r.GetProductTx(tx, id)
	if err != nil {
		// span.RecordError(err)
		return 0, 0, err
	}

	return pdt.Price, pdt.Price * float64(qty), nil
}

func (r productRepo) GetProductTx(tx *gorm.DB, id uint64) (*entity.Product, error) {
	var pdt *entity.Product

	cacheRepo := cache.NewCacheRepository(r.p, "redis")
	_ = cacheRepo.GetKey(fmt.Sprintf("%s%d", redis_entity.RedisProductData, id), &pdt)

	if pdt == nil {
		err := tx.Debug().Unscoped().Preload("Inventory").Where("product_id = ?", id).Take(&pdt).Error
		// err := r.p.ProductDb.Debug().Where("product_id = ?", id).Take(&pdt).Error
		if err != nil {
			return nil, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		_ = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisProductData, id), pdt, redis_entity.RedisExpirationGlobal)
	}

	return pdt, nil
}
