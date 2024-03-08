package warehouse

import (
	"errors"
	"log"
	entity "products-crud/domain/entity/warehouse_entity"
	repository "products-crud/domain/repository/warehouse_repository"

	base "products-crud/infrastructure/persistences"

	"gorm.io/gorm"
)

type warehouseRepo struct {
	p *base.Persistence
}

func NewWarehouseRepository(p *base.Persistence) repository.WarehouseRepository {
	return &warehouseRepo{p}
}

func (r warehouseRepo) AddWarehouse(w *entity.Warehouse) (*entity.Warehouse, error) {
	log.Println("Adding new warehouse ", w.Name, "...")

	if err := r.p.ProductDb.Debug().Create(&w).Error; err != nil {
		return nil, err
	}

	log.Println(w.Name, " created.")
	return w, nil
}

func (r warehouseRepo) GetWarehouse(id uint64) (*entity.Warehouse, error) {
	var w *entity.Warehouse

	err := r.p.ProductDb.Debug().Unscoped().Where("warehouse_id = ?", id).Take(&w).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("warehouse not found")
	}

	return w, nil
}

func (r warehouseRepo) GetWarehouses() ([]entity.Warehouse, error) {
	var w []entity.Warehouse

	err := r.p.ProductDb.Debug().Find(&w).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("warehouse not found")
	}

	return w, nil
}

func (r warehouseRepo) UpdateWarehouse(w *entity.Warehouse) (*entity.Warehouse, error) {
	err := r.p.ProductDb.Debug().Where("warehouse_id = ?", w.WarehouseID).Updates(&w).Error

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (r warehouseRepo) DeleteWarehouse(id uint64) error {
	var w entity.Warehouse
	res := r.p.ProductDb.Debug().Where("warehouse_id = ?", id).Delete(&w)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("warehouse not found")
	}

	return nil
}
