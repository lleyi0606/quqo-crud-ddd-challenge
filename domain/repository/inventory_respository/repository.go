package repository

import (
	entity "products-crud/domain/entity/inventory_entity"
)

type InventoryRepository interface {
	AddInventory(*entity.Inventory) (*entity.Inventory, error)
	GetInventory(uint64) (*entity.Inventory, error)
	GetInventories() ([]entity.Inventory, error)
	UpdateInventory(*entity.Inventory) (*entity.Inventory, error)
	DeleteInventory(uint64) (*entity.Inventory, error)
	SearchInventory(string) ([]entity.Inventory, error)
}

type InventoryHandlerRepository interface {
	AddInventory(*entity.Inventory) (*entity.Inventory, error)
	GetInventory(uint64) (*entity.Inventory, error)
	GetInventories() ([]entity.Inventory, error)
	UpdateInventory(*entity.Inventory) (*entity.Inventory, error)
	DeleteInventory(uint64) (*entity.Inventory, error)
	SearchInventory(string) ([]entity.Inventory, error)
}
