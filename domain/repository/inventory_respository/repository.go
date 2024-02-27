package repository

import (
	entity "products-crud/domain/entity/inventory_entity"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	AddInventory(*entity.Inventory) (*entity.Inventory, error)
	GetInventory(uint64) (*entity.Inventory, error)
	GetInventoryTx(*gorm.DB, uint64) (*entity.Inventory, error)
	// GetInventories() ([]entity.Inventory, error)
	UpdateInventory(*entity.Inventory) (*entity.Inventory, error)
	UpdateInventoryTx(*gorm.DB, *entity.Inventory) (*entity.Inventory, error)
	DeleteInventory(uint64) (*entity.Inventory, error)
	// HasSufficientStock(uint64, int) error
	DecreaseStock(uint64, int) error
	DecreaseStockTx(*gorm.DB, uint64, int) error

	// SearchInventory(string) ([]entity.Inventory, error)
}

type InventoryHandlerRepository interface {
	// AddInventory(*entity.Inventory) (*entity.Inventory, error)
	// GetInventory(uint64) (*entity.Inventory, error)
	// GetInventories() ([]entity.Inventory, error)
	// UpdateInventory(*entity.Inventory) (*entity.Inventory, error)
	// DeleteInventory(uint64) (*entity.Inventory, error)
	// SearchInventory(string) ([]entity.Inventory, error)

	GetInventory(uint64) (*entity.Inventory, error)
	UpdateStock(uint64, *entity.InventoryStockOnly) (*entity.Inventory, error)
}
