package inventory

import (
	"errors"
	"fmt"
	"log"
	entity "products-crud/domain/entity/inventory_entity"
	"products-crud/domain/entity/redis_entity"
	repository "products-crud/domain/repository/inventory_respository"
	"products-crud/infrastructure/implementations/cache"
	base "products-crud/infrastructure/persistences"

	"gorm.io/gorm"
)

type inventoryRepo struct {
	p *base.Persistence
}

func NewInventoryRepository(p *base.Persistence) repository.InventoryRepository {
	return &inventoryRepo{p}
}

func (r inventoryRepo) AddInventory(ivt *entity.Inventory) (*entity.Inventory, error) {
	log.Println("Adding new inventory ", ivt.ProductID, "...")

	if err := r.p.ProductDb.Debug().Create(&ivt).Error; err != nil {
		return nil, err
	}

	// add to search repo
	// searchRepo := search.NewSearchRepository(r.p, "algolia")
	// err := searchRepo.AddInventory(ivt)
	// if err != nil {
	// 	return nil, err
	// }

	log.Println(ivt.ProductID, " created.")
	return ivt, nil
}

func (r inventoryRepo) GetInventory(id uint64) (*entity.Inventory, error) {
	var ivt *entity.Inventory

	cacheRepo := cache.NewCacheRepository(r.p, "redis")
	_ = cacheRepo.GetKey(fmt.Sprintf("%s%d", redis_entity.RedisInventoryData, id), &ivt)

	if ivt == nil {
		log.Print("inventory not found in redis")
		err := r.p.ProductDb.Debug().Where("product_id = ?", id).Take(&ivt).Error
		if err != nil {
			return nil, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found")
		}

		_ = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisInventoryData, id), ivt, redis_entity.RedisExpirationGlobal)
	}

	return ivt, nil
}

// func (r inventoryRepo) GetInventories() ([]entity.Inventory, error) {
// 	var ivts []entity.Inventory
// 	err := r.p.ProductDb.Debug().Find(&ivts).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, errors.New("inventory not found")
// 	}
// 	return ivts, nil
// }

func (r inventoryRepo) UpdateInventory(ivt *entity.Inventory) (*entity.Inventory, error) {
	// err := r.p.ProductDb.Debug().Model(&entity.Inventory{}).Where("id = ?", ivt.ID).Updates(ivt).Error
	err := r.p.ProductDb.Debug().Where("product_id = ?", ivt.ProductID).Updates(&ivt).Error

	if err != nil {
		return nil, err
	}

	// update cache
	cacheRepo := cache.NewCacheRepository(r.p, "redis")
	err = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisInventoryData, ivt.ProductID), &ivt, redis_entity.RedisExpirationGlobal)
	if err != nil {
		return nil, err
	}

	// update search repo
	// searchRepo := search.NewSearchRepository(r.p, "algolia")
	// err = searchRepo.UpdateInventory(ivt)
	// if err != nil {
	// 	log.Print(err)
	// 	return nil, err
	// }

	return ivt, nil
}

func (r inventoryRepo) DeleteInventory(id uint64) (*entity.Inventory, error) {
	var ivt entity.Inventory
	err := r.p.ProductDb.Debug().Where("product_id = ?", id).Delete(&ivt).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("inventory not found")
	}

	// search repo
	// searchRepo := search.NewSearchRepository(r.p, "algolia")
	// err = searchRepo.DeleteInventory(id)
	// if err != nil {
	// 	return nil, err
	// }

	// update cache
	cacheRepo := cache.NewCacheRepository(r.p, "redis")
	err = cacheRepo.DeleteRecord(fmt.Sprintf("%s%d", redis_entity.RedisInventoryData, id))
	if err != nil {
		return nil, err
	}

	return &ivt, nil
}

// func (r inventoryRepo) SearchInventory(str string) ([]entity.Inventory, error) {

// 	// new search repo
// 	searchRepo := search.NewSearchRepository(r.p, "algolia")
// 	ivts, err := searchRepo.SearchInventories(str)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ivts, nil
// }
