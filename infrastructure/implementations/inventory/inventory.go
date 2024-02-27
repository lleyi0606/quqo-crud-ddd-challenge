package inventory

import (
	"context"
	"errors"
	"fmt"
	"log"
	entity "products-crud/domain/entity/inventory_entity"
	"products-crud/domain/entity/redis_entity"
	repository "products-crud/domain/repository/inventory_respository"
	"products-crud/infrastructure/implementations/cache"
	base "products-crud/infrastructure/persistences"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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

		// _ = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisInventoryData, id), ivt, redis_entity.RedisExpirationGlobal)
	}

	return ivt, nil
}

func (r inventoryRepo) GetInventoryTx(tx *gorm.DB, id uint64) (*entity.Inventory, error) {
	var ivt *entity.Inventory

	if tx == nil {
		tx = r.p.ProductDb.Begin()
		var errTx error
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			} else if errTx != nil {
				tx.Rollback()
			} else {
				errC := tx.Commit().Error
				if errC != nil {
					tx.Rollback()
				}
			}
		}()
	}

	if ivt == nil {
		log.Print("inventory not found in redis")
		err := tx.Debug().Where("product_id = ?", id).Take(&ivt).Error
		if err != nil {
			return nil, err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inventory not found")
		}

		// _ = cacheRepo.SetKey(fmt.Sprintf("%s%d", redis_entity.RedisInventoryData, id), ivt, redis_entity.RedisExpirationGlobal)
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

	err = cacheRepo.DeleteRecord(fmt.Sprintf("%s%d", redis_entity.RedisProductData, ivt.ProductID))
	if err != nil {
		return nil, err
	}

	return ivt, nil
}

func (r inventoryRepo) UpdateInventoryTx(tx *gorm.DB, ivt *entity.Inventory) (*entity.Inventory, error) {
	err := tx.Debug().Where("product_id = ?", ivt.ProductID).Updates(&ivt).Error

	if err != nil {
		return nil, err
	}

	// update cache
	cacheRepo := cache.NewCacheRepository(r.p, "redis")

	err = cacheRepo.DeleteRecord(fmt.Sprintf("%s%d", redis_entity.RedisProductData, ivt.ProductID))
	if err != nil {
		return nil, err
	}

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

// func (r inventoryRepo) HasSufficientStock(id uint64, stock int) error {
// 	ivt, err := r.GetInventory(id)
// 	if err != nil {
// 		return err
// 	}
// 	if ivt.Stock < stock {
// 		return fmt.Errorf("insufficient stock for product_id %d", id)
// 	}
// 	return nil
// }

func (r inventoryRepo) DecreaseStock(id uint64, qty int) error {
	ivt, err := r.GetInventory(id)
	if err != nil {
		return err
	}
	if ivt.Stock < qty {
		return fmt.Errorf("insufficient stock for product_id %d", id)
	}

	ivt.Stock -= qty

	_, err = r.UpdateInventory(ivt)

	return err
}

func (r inventoryRepo) DecreaseStockTx(tx *gorm.DB, id uint64, qty int, ctx context.Context) error {

	tracer := otel.Tracer("quqo")

	// Start a new span for the function
	_, span := tracer.Start(ctx, "implementation/DecreaseStockTx",
		trace.WithAttributes(
			attribute.String("Description", "DecreaseStockTx in implementation"),
		),
	)
	defer span.End()

	var stock int
	err := tx.Raw("SELECT stock FROM inventories WHERE product_id = ?", id).Scan(&stock)
	if err.Error != nil {
		span.RecordError(err.Error)
		return err.Error
	}
	if stock < qty {
		stockErr := fmt.Errorf("insufficient stock for product_id %d", id)
		span.RecordError(stockErr)
		return stockErr
	}

	err = tx.Exec("UPDATE inventories SET stock = ? WHERE product_id = ?", gorm.Expr("stock - ?", qty), id)

	return err.Error

}
