package search_repository

import (
	inventory_entity "products-crud/domain/entity/inventory_entity"
	entity "products-crud/domain/entity/product_entity"
)

type SearchRepository interface {
	AddProduct(*entity.Product) error
	SearchProducts(string) ([]entity.Product, error)
	DeleteProduct(uint64) error
	UpdateProduct(*entity.Product) error

	AddInventory(*inventory_entity.Inventory) error
	SearchInventories(string) ([]inventory_entity.Inventory, error)
	DeleteInventory(uint64) error
	UpdateInventory(*inventory_entity.Inventory) error
}
