package entity

import (
	"products-crud/domain/entity"
)

type OrderedItem struct {
	entity.BaseModelWDelete
	OrderID    uint64  `gorm:"type:numeric;" json:"order_id"`
	ProductID  string  `gorm:"type:numeric;" json:"product_id"`
	Quantity   int     `gorm:"type:numeric;" json:"quantity"`
	UnitPrice  float64 `gorm:"type:numeric;" json:"unit_price"`
	TotalPrice float64 `gorm:"type:numeric;" json:"total_price"`
}

type OrderedItemInput struct {
	entity.BaseModelWDelete
	ProductID string `gorm:"type:numeric;" json:"product_id"`
	Quantity  int    `gorm:"type:numeric;" json:"quantity"`
}
