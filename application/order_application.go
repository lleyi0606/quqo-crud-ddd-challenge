package application

import (
	entity "products-crud/domain/entity/order_entity"
	repository "products-crud/domain/repository/order_repository"
	"products-crud/infrastructure/implementations/order"
	base "products-crud/infrastructure/persistences"
)

type OrderApp struct {
	p *base.Persistence
}

func NewOrderApplication(p *base.Persistence) repository.OrderHandlerRepository {
	return &OrderApp{p}
}

func (u *OrderApp) AddOrder(cat *entity.Order) (*entity.Order, error) {
	repoOrder := order.NewOrderRepository(u.p)
	return repoOrder.AddOrder(cat)
}

func (u *OrderApp) GetOrder(id uint64) (*entity.Order, error) {
	repoOrder := order.NewOrderRepository(u.p)
	return repoOrder.GetOrder(id)
}

func (u *OrderApp) UpdateOrder(cat *entity.Order) (*entity.Order, error) {
	repoOrder := order.NewOrderRepository(u.p)
	return repoOrder.UpdateOrder(cat)
}

func (u *OrderApp) DeleteOrder(id uint64) error {
	repoOrder := order.NewOrderRepository(u.p)
	return repoOrder.DeleteOrder(id)
}
