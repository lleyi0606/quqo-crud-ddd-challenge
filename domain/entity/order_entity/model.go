package entity

import (
	"products-crud/domain/entity"
)

type Order struct {
	entity.BaseModelWDelete
	OrderID       uint64  `gorm:"primary_key;auto_increment:true" json:"order_id"`
	TotalCost     float64 `gorm:"type:numeric;" json:"total_cost"`
	TotalFees     float64 `gorm:"type:numeric;" json:"total_fees"`
	TotalCheckout float64 `gorm:"type:numeric;" json:"total_checkout"`
	CustomerID    uint64  `gorm:"type:numeric;" json:"customer_id"`
	WarehouseID   uint64  `gorm:"type:numeric;" json:"warehouse_id"`
	Status        string  `gorm:"size:100;" json:"status"`

	// Customer customer_entity.Customer `gorm:"foreignkey:CustomerID" json:"customer"`
}
