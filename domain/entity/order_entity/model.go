package entity

import (
	"products-crud/domain/entity"
	orderItem_entity "products-crud/domain/entity/orderedItem_entity"
)

type Order struct {
	entity.BaseModelWDelete
	OrderID       uint64                         `gorm:"primary_key;auto_increment:true" json:"order_id"`
	CustomerID    uint64                         `gorm:"type:numeric;" json:"customer_id"`
	WarehouseID   uint64                         `gorm:"type:numeric;" json:"warehouse_id"`
	Status        string                         `gorm:"size:100;" json:"status"`
	OrderedItems  []orderItem_entity.OrderedItem `gorm:"foreignKey:OrderID" json:"ordered_items"`
	TotalCost     float64                        `gorm:"type:numeric;" json:"total_cost"`
	TotalFees     float64                        `gorm:"type:numeric;" json:"total_fees"`
	TotalCheckout float64                        `gorm:"type:numeric;" json:"total_checkout"`
}

type OrderInput struct {
	entity.BaseModelWDelete
	OrderID      uint64                              `gorm:"primary_key;auto_increment:true" json:"order_id"`
	CustomerID   uint64                              `gorm:"type:numeric;" json:"customer_id"`
	WarehouseID  uint64                              `gorm:"type:numeric;" json:"warehouse_id"`
	Status       string                              `gorm:"size:100;" json:"status"`
	OrderedItems []orderItem_entity.OrderedItemInput `gorm:"foreignKey:OrderID" json:"ordered_items"`
}
