package repository

import entity "products-crud/domain/entity/orderedItem_entity"

type OrderedItemRepository interface {
	AddOrderedItem(*entity.OrderedItem) (*entity.OrderedItem, error)
	GetOrderedItems() ([]entity.OrderedItem, error)
	GetOrderedItemsByOrderId(uint64) ([]entity.OrderedItem, error)
}

type OrderedItemHandlerRepository interface {
	GetOrderedItems() ([]entity.OrderedItem, error)
	GetOrderedItemsByOrderId(uint64) ([]entity.OrderedItem, error)
}
