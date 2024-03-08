package entity

import "products-crud/domain/entity"

type Inventory struct {
	entity.BaseModelWDelete
	ProductID   string `gorm:"primaryKey;autoIncrement:true" json:"product_id"`
	WarehouseID int64  `gorm:"type:numeric" json:"warehouse_id"`
	Stock       int    `gorm:"type:numeric;" json:"stock"`
}

type InventoryUser struct {
	WarehouseID int64 `gorm:"type:numeric" json:"warehouse_id"`
	Stock       int   `gorm:"type:numeric;" json:"stock"`
}

type InventoryAlgolia struct {
	Inventory
	ObjectID string `json:"object_id"`
}

type InventoryStockOnly struct {
	Stock int `gorm:"type:numeric;" json:"stock"`
}
