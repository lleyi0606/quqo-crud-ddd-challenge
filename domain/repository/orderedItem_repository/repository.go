package repository

import (
	entity "products-crud/domain/entity/orderedItem_entity"
)

type OrderedItemRepository interface {
	GetOrderedItems() ([]entity.OrderedItem, error)
	GetOrderedItem(uint64) (*entity.OrderedItem, error)
	UpdateOrderedItem(*entity.OrderedItem) (*entity.OrderedItem, error)
	DeleteOrderedItem(uint64) error
}

type OrderedItemHandlerRepository interface {
	GetOrderedItems() ([]entity.OrderedItem, error)
	GetOrderedItem(uint64) (*entity.OrderedItem, error)
	UpdateOrderedItem(*entity.OrderedItem) (*entity.OrderedItem, error)
	DeleteOrderedItem(uint64) error
}
