package entity

import (
	"fmt"
	"math/rand"
	"products-crud/domain/entity"
	inventory_entity "products-crud/domain/entity/inventory_entity"
	"time"
)

type ProductWithStockAndWarehouse struct {
	entity.BaseModelWDelete
	ProductID   string  `gorm:"primary_key" json:"product_id"`
	Name        string  `gorm:"size:100;" json:"name"`
	Description string  `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"type:numeric;" json:"price"`
	Status      string  `gorm:"size:100;" json:"status"`
	CategoryID  int64   `gorm:"type:numeric;" json:"category_id"`
	WarehouseID int64   `gorm:"type:numeric" json:"warehouse_id"`
	Stock       int     `gorm:"type:numeric;" json:"stock"`
}

type Product struct {
	entity.BaseModelWDelete
	ProductID   string                     `gorm:"primary_key" json:"product_id"`
	Name        string                     `gorm:"size:100;" json:"name"`
	Description string                     `gorm:"size:255;" json:"description"`
	Price       float64                    `gorm:"type:numeric;" json:"price"`
	Status      string                     `gorm:"size:100;" json:"status"`
	CategoryID  int64                      `gorm:"type:numeric;" json:"category_id"`
	Inventory   inventory_entity.Inventory `gorm:"foreignkey:ProductID;references:ProductID" json:"inventory"`
}

type ProductSearch struct {
	Name string `json:"name"`
	// ObjectID  int    `json:"object_id"`
	ProductID string `json:"product_id"`
}

func GenerateProductID() string {
	// You can replace this with your logic to generate a unique numeric identifier.
	uniqueNumber := generateUniqueNumber()

	return fmt.Sprintf("PDT-%06d", uniqueNumber)
}

// generateUniqueNumber generates a random unique number.
func generateUniqueNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(1000000)
}
