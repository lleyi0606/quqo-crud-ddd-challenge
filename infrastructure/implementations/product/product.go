package product

import (
	"errors"
	"fmt"
	"log"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/entity/redis_entity"
	"products-crud/infrastructure/implementations/cache"
	"products-crud/infrastructure/implementations/search"
	base "products-crud/infrastructure/persistences"

	"gorm.io/gorm"
)

type productRepo struct {
	p *base.Persistence
}

func NewProductRepository(p *base.Persistence) *productRepo {
	return &productRepo{p}
}

// productRepo implements the repository.ProductRepository interface

func (r productRepo) AddProduct(pdt *entity.Product) (*entity.Product, error) {
	log.Println("Adding new product ", pdt.Name, "...")

	if err := r.p.ProductDb.Debug().Create(&pdt).Error; err != nil {
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
		err := r.p.ProductDb.Debug().Where("id = ?", id).Take(&pdt).Error
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

func (r productRepo) UpdateProduct(pdt *entity.ProductUpdate) (*entity.Product, error) {
	err := r.p.ProductDb.Debug().Model(&entity.Product{}).Where("id = ?", pdt.ID).Updates(pdt).Error

	if err != nil {
		return nil, err
	}

	return r.GetProduct(pdt.ID)
}

func (r productRepo) DeleteProduct(id uint64) (*entity.Product, error) {
	var pdt entity.Product
	err := r.p.ProductDb.Debug().Where("id = ?", id).Delete(&pdt).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
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

// Search from CockroachDB
// func (r productRepo) SearchProductsCockroach(str string) ([]entity.Product, error) {
// 	var pdts []entity.Product

// 	// cacheRepo := redis.NewRedisRepository(r.p)
// 	// _ = cacheRepo.SearchName(str, &pdts)

// 	// if pdts == nil {

// 	err := r.p.ProductDb.Debug().Where("lower(name) LIKE lower(?)", "%"+str+"%").Find(&pdts).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, errors.New("product not found")
// 	}

// 	// var pdt []entity.Product

// 	// _ = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisProductData, id), pdt, redis_entity.RedisExpirationGlobal)

// 	// }

// 	return pdts, nil
// }
