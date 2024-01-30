package entity

import (
	entity "products-crud/domain/entity/inventory_entity"
)

type ProductWithoutInventory struct {
	ProductID   uint64  `gorm:"primary_key;auto_increment" json:"product_id"`
	Name        string  `gorm:"size:100;" json:"name"`
	Description *string `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"type:numeric;" json:"price"`
	Category    string  `gorm:"size:100;" json:"category"`
}

type ProductWithStockAndWarehouse struct {
	ProductID   uint64  `gorm:"primary_key;auto_increment" json:"product_id"`
	Name        string  `gorm:"size:100;" json:"name"`
	Description string  `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"type:numeric;" json:"price"`
	Category    string  `gorm:"size:100;" json:"category"`
	WarehouseID int64   `gorm:"type:numeric" json:"warehouse_id"`
	Stock       int     `gorm:"type:numeric;" json:"stock"`
}

type Product struct {
	ProductID   uint64           `gorm:"primary_key;auto_increment:false" json:"product_id"`
	Name        string           `gorm:"size:100;" json:"name"`
	Description string           `gorm:"size:255;" json:"description"`
	Price       float64          `gorm:"type:numeric;" json:"price"`
	Category    string           `gorm:"size:100;" json:"category"`
	Inventory   entity.Inventory `gorm:"foreignkey:ProductID;references:ProductID" json:"inventory"`
}

type ProductAlgolia struct {
	Product
	ObjectID uint64 `json:"object_id"`
}
