package repository

import (
	entity "products-crud/domain/entity/warehouse_entity"
)

type WarehouseRepository interface {
	AddWarehouse(*entity.Warehouse) (*entity.Warehouse, error)
	GetWarehouse(uint64) (*entity.Warehouse, error)
	GetWarehouses() ([]entity.Warehouse, error)
	UpdateWarehouse(*entity.Warehouse) (*entity.Warehouse, error)
	DeleteWarehouse(uint64) error
}

type WarehouseHandlerRepository interface {
	AddWarehouse(*entity.Warehouse) (*entity.Warehouse, error)
	GetWarehouse(uint64) (*entity.Warehouse, error)
	GetWarehouses() ([]entity.Warehouse, error)
	UpdateWarehouse(*entity.Warehouse) (*entity.Warehouse, error)
	DeleteWarehouse(uint64) error
}
