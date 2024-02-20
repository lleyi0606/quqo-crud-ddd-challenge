package repository

import (
	entity "products-crud/domain/entity/order_entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	AddOrder(*entity.Order) (*entity.Order, error)
	AddOrderTx(*gorm.DB, *entity.Order) (*entity.Order, error)
	GetOrder(uint64) (*entity.Order, error)
	UpdateOrder(*entity.Order) (*entity.Order, error)
	DeleteOrder(uint64) error
}

type OrderHandlerRepository interface {
	AddOrder(*entity.OrderInput) (*entity.Order, error)
	GetOrder(uint64) (*entity.Order, error)
	UpdateOrder(*entity.Order) (*entity.Order, error)
	DeleteOrder(uint64) error
}
