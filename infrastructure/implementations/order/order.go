package order

import (
	"context"
	"errors"
	entity "products-crud/domain/entity/order_entity"
	repository "products-crud/domain/repository/order_repository"

	base "products-crud/infrastructure/persistences"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type orderRepo struct {
	p *base.Persistence
	c *context.Context
}

func NewOrderRepository(p *base.Persistence, c *context.Context) repository.OrderRepository {
	return &orderRepo{p, c}
}

func (r orderRepo) AddOrder(order *entity.Order) (*entity.Order, error) {

	if err := r.p.ProductDb.Debug().Create(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r orderRepo) AddOrderTx(tx *gorm.DB, order *entity.Order) (*entity.Order, error) {

	tracer := otel.Tracer("quqo")

	// Start a new span for the function
	_, span := tracer.Start(*r.c, "implementation/AddOrderTx",
		trace.WithAttributes(
			attribute.String("Description", "AddOrderTx in implementation"),
		),
	)
	defer span.End()

	if err := tx.Debug().Create(&order).Error; err != nil {
		span.RecordError(err)
		return nil, err
	}

	return order, nil
}

func (r orderRepo) GetOrder(id uint64) (*entity.Order, error) {
	var order *entity.Order

	err := r.p.ProductDb.Debug().Unscoped().Preload("OrderedItems").Where("order_id = ?", id).Take(&order).Error

	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r orderRepo) UpdateOrder(cus *entity.Order) (*entity.Order, error) {
	result := r.p.ProductDb.Debug().Where("order_id = ?", cus.OrderID).Updates(&cus)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("order not found")
	}

	return cus, nil
}

func (r orderRepo) DeleteOrder(id uint64) error {
	var order entity.Order
	res := r.p.ProductDb.Debug().Where("order_id = ?", id).Delete(&order)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}
