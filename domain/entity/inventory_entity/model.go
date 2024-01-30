package entity

type Inventory struct {
	ProductID   uint64 `gorm:"primaryKey;autoIncrement:true" json:"product_id"`
	WarehouseID int64  `gorm:"type:numeric" json:"warehouse_id"`
	Stock       int    `gorm:"type:numeric;" json:"stock"`
}

type InventoryAlgolia struct {
	Inventory
	ObjectID uint64 `json:"object_id"`
}

type InventoryStockOnly struct {
	Stock int `gorm:"type:numeric;" json:"stock"`
}
