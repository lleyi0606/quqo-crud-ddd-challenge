package application

import (
	entity "products-crud/domain/entity/orderedItem_entity"
	repository "products-crud/domain/repository/orderedItem_repository"
	"products-crud/infrastructure/implementations/orderedItem"
	base "products-crud/infrastructure/persistences"
)

type OrderedItemApp struct {
	p *base.Persistence
}

func NewOrderedItemApplication(p *base.Persistence) repository.OrderedItemHandlerRepository {
	return &OrderedItemApp{p}
}

func (u *OrderedItemApp) GetOrderedItems() ([]entity.OrderedItem, error) {
	repoOrderedItem := orderedItem.NewOrderedItemRepository(u.p, nil)
	return repoOrderedItem.GetOrderedItems()
}

func (u *OrderedItemApp) GetOrderedItemsByOrderId(id uint64) ([]entity.OrderedItem, error) {
	repoOrderedItem := orderedItem.NewOrderedItemRepository(u.p, nil)
	return repoOrderedItem.GetOrderedItemsByOrderId(id)
}
