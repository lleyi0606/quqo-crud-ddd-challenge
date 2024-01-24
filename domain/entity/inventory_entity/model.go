package entity

type Inventory struct {
	InventoryID uint64 `gorm:"primary_key;auto_increment" json:"i_id"`
	Name        string `gorm:"size:100;" json:"name"`
	Address     string `gorm:"size:255;" json:"address"`
}

type InventoryAlgolia struct {
	Inventory
	ObjectID uint64 `json:"objectID"`
}
