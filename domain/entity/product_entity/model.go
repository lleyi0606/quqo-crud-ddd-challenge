package entity

type Product struct {
	ID          uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Name        string  `gorm:"size:100;" json:"name"`
	Description *string `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"type:numeric;" json:"price"`
	Category    string  `gorm:"size:100;" json:"category"`
	Stock       int     `gorm:"type:numeric;" json:"stock"`
	Image       string  `gorm:"size:255;" json:"image"`
	InventoryID uint64  `gorm:"foreign_key" json:"i_id"`
}

type ProductAlgolia struct {
	// ID          uint64  `json:"id"`
	// Name        string  `json:"name"`
	// Description string  `json:"description"`
	// Price       float64 `json:"price"`
	// Category    string  `json:"category"`
	// Stock       int     `json:"stock"`
	// Image       string  `json:"image"`
	// ProductUpdate
	Product
	ObjectID uint64 `json:"objectID"`
}
