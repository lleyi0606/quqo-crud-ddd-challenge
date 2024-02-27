package repository

import (
	"context"
	entity "products-crud/domain/entity/order_entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	AddOrder(*entity.Order) (*entity.Order, error)
	AddOrderTx(*gorm.DB, *entity.Order, context.Context) (*entity.Order, error)
	GetOrder(uint64) (*entity.Order, error)
	UpdateOrder(*entity.Order) (*entity.Order, error)
	DeleteOrder(uint64) error
}

type OrderHandlerRepository interface {
	AddOrder(*entity.OrderInput, context.Context) (*entity.Order, error)
	GetOrder(uint64) (*entity.Order, error)
	UpdateOrder(*entity.Order) (*entity.Order, error)
	DeleteOrder(uint64) error
}
