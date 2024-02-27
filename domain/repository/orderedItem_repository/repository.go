package repository

import (
	"context"
	entity "products-crud/domain/entity/orderedItem_entity"

	"gorm.io/gorm"
)

type OrderedItemRepository interface {
	AddOrderedItem(*entity.OrderedItem) (*entity.OrderedItem, error)
	AddOrderedItemTx(*gorm.DB, *entity.OrderedItem, context.Context) (*entity.OrderedItem, error)
	GetOrderedItems() ([]entity.OrderedItem, error)
	GetOrderedItemsByOrderId(uint64) ([]entity.OrderedItem, error)
}

type OrderedItemHandlerRepository interface {
	GetOrderedItems() ([]entity.OrderedItem, error)
	GetOrderedItemsByOrderId(uint64) ([]entity.OrderedItem, error)
}
