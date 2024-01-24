package application

import (
	entity "products-crud/domain/entity/inventory_entity"
	repository "products-crud/domain/repository/inventory_respository"
	"products-crud/infrastructure/implementations/inventory"
	base "products-crud/infrastructure/persistences"
)

type inventoryApp struct {
	p *base.Persistence
}

func NewInventoryApplication(p *base.Persistence) repository.InventoryHandlerRepository {
	return &inventoryApp{p}
}

func (u *inventoryApp) AddInventory(user *entity.Inventory) (*entity.Inventory, error) {
	repoInventory := inventory.NewInventoryRepository(u.p)
	return repoInventory.AddInventory(user)
}

func (u *inventoryApp) GetInventory(ivtId uint64) (*entity.Inventory, error) {
	repoInventory := inventory.NewInventoryRepository(u.p)
	return repoInventory.GetInventory(ivtId)
}

func (u *inventoryApp) GetInventories() ([]entity.Inventory, error) {
	repoInventory := inventory.NewInventoryRepository(u.p)
	return repoInventory.GetInventories()
}

func (u *inventoryApp) UpdateInventory(pdt *entity.Inventory) (*entity.Inventory, error) {
	repoInventory := inventory.NewInventoryRepository(u.p)
	return repoInventory.UpdateInventory(pdt)
}

func (u *inventoryApp) DeleteInventory(ivtId uint64) (*entity.Inventory, error) {
	// repoProduct := product.NewProductRepository(u.p)

	repoInventory := inventory.NewInventoryRepository(u.p)
	return repoInventory.DeleteInventory(ivtId)
}

func (u *inventoryApp) SearchInventory(str string) ([]entity.Inventory, error) {
	repoInventory := inventory.NewInventoryRepository(u.p)
	return repoInventory.SearchInventory(str)
}
