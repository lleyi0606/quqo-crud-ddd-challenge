package entity

type Inventory struct {
	ProductID   uint64 `gorm:"primary_key;auto_increment" json:"i_id"`
	WarehouseID uint64 `gorm:"type:numeric" json:"warehouse_id"`
	Stock       int    `gorm:"type:numeric;" json:"stock"`
}

type InventoryAlgolia struct {
	Inventory
	ObjectID uint64 `json:"objectID"`
}

type InventoryStockOnly struct {
	Stock int `gorm:"type:numeric;" json:"stock"`
}
