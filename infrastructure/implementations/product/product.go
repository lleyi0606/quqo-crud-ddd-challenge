package product

import (
	"errors"
	"fmt"
	"log"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/entity/redis_entity"
	repository "products-crud/domain/repository/product_respository"
	"products-crud/infrastructure/implementations/cache"
	"products-crud/infrastructure/implementations/inventory"
	"products-crud/infrastructure/implementations/search"
	base "products-crud/infrastructure/persistences"

	"gorm.io/gorm"
)

type productRepo struct {
	p *base.Persistence
}

func NewProductRepository(p *base.Persistence) repository.ProductRepository {
	return &productRepo{p}
}

// productRepo implements the repository.ProductRepository interface

func (r productRepo) AddProduct(pdt *entity.Product) (*entity.Product, error) {
	log.Println("Adding new product ", pdt.Name, "...")

	if err := r.p.ProductDb.Debug().Create(&pdt).Error; err != nil {
		return nil, err
	}

	// add to search repo
	searchRepo := search.NewSearchRepository(r.p, "opensearch")
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
		err := r.p.ProductDb.Debug().Preload("Inventory").Where("product_id = ?", id).Take(&pdt).Error
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
	searchRepo := search.NewSearchRepository(r.p, "opensearch")
	err = searchRepo.UpdateProduct(pdt)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return pdt, nil
}

func (r productRepo) DeleteProduct(id uint64) (*entity.Product, error) {
	var pdt entity.Product
	err := r.p.ProductDb.Debug().Where("product_id = ?", id).Delete(&pdt).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}

	// delete from inventory too
	inventoryRepo := inventory.NewInventoryRepository(r.p)
	_, err = inventoryRepo.DeleteInventory(id)
	if err != nil {
		return nil, err
	}

	// search repo
	searchRepo := search.NewSearchRepository(r.p, "opensearch")
	err = searchRepo.DeleteProduct(id)
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
	searchRepo := search.NewSearchRepository(r.p, "opensearch")
	pdts, err := searchRepo.SearchProducts(str)
	if err != nil {
		return nil, err
	}
	return pdts, nil

}
