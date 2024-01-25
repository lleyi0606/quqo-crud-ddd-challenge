package entity

type Inventory struct {
	ProductID   uint64 `gorm:"primaryKey;autoIncrement:true" json:"productId"`
	WarehouseID uint64 `gorm:"type:numeric" json:"warehouseId"`
	Stock       int    `gorm:"type:numeric;" json:"stock"`
}

type InventoryAlgolia struct {
	Inventory
	ObjectID uint64 `json:"objectId"`
}

type InventoryStockOnly struct {
	Stock int `gorm:"type:numeric;" json:"stock"`
}
