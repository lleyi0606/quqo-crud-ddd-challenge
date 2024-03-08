package application

import (
	entity "products-crud/domain/entity/warehouse_entity"
	repository "products-crud/domain/repository/warehouse_repository"
	"products-crud/infrastructure/implementations/warehouse"
	base "products-crud/infrastructure/persistences"
)

type WarehouseApp struct {
	p *base.Persistence
}

func NewWarehouseApplication(p *base.Persistence) repository.WarehouseHandlerRepository {
	return &WarehouseApp{p}
}

func (u *WarehouseApp) AddWarehouse(cat *entity.Warehouse) (*entity.Warehouse, error) {
	repoWarehouse := warehouse.NewWarehouseRepository(u.p)
	return repoWarehouse.AddWarehouse(cat)
}

func (u *WarehouseApp) GetWarehouse(id uint64) (*entity.Warehouse, error) {
	repoWarehouse := warehouse.NewWarehouseRepository(u.p)
	return repoWarehouse.GetWarehouse(id)
}

func (u *WarehouseApp) GetWarehouses() ([]entity.Warehouse, error) {
	repoWarehouse := warehouse.NewWarehouseRepository(u.p)
	return repoWarehouse.GetWarehouses()
}

func (u *WarehouseApp) UpdateWarehouse(cat *entity.Warehouse) (*entity.Warehouse, error) {
	repoWarehouse := warehouse.NewWarehouseRepository(u.p)
	return repoWarehouse.UpdateWarehouse(cat)
}

func (u *WarehouseApp) DeleteWarehouse(id uint64) error {
	repoWarehouse := warehouse.NewWarehouseRepository(u.p)
	return repoWarehouse.DeleteWarehouse(id)
}
