package entity

import (
	entity "products-crud/domain/entity/inventory_entity"
)

type ProductWithoutInventory struct {
	ProductID   uint64  `gorm:"primary_key;auto_increment" json:"productId"`
	Name        string  `gorm:"size:100;" json:"name"`
	Description *string `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"type:numeric;" json:"price"`
	Category    string  `gorm:"size:100;" json:"category"`
}

type ProductWithStockAndWarehouse struct {
	ProductID   uint64  `gorm:"primary_key;auto_increment" json:"productId"`
	Name        string  `gorm:"size:100;" json:"name"`
	Description string  `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"type:numeric;" json:"price"`
	Category    string  `gorm:"size:100;" json:"category"`
	WarehouseID uint64  `gorm:"type:numeric" json:"warehouseId"`
	Stock       int     `gorm:"type:numeric;" json:"stock"`
}

type Product struct {
	ProductID   uint64  `gorm:"primary_key;auto_increment:false" json:"productId"`
	Name        string  `gorm:"size:100;" json:"name"`
	Description string  `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"type:numeric;" json:"price"`
	Category    string  `gorm:"size:100;" json:"category"`
	Inventory   entity.Inventory
}

type ProductAlgolia struct {
	Product
	ObjectID uint64 `json:"objectId"`
}
