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

	tracer := otel.Tracer("quqo")
	_, span := tracer.Start(*r.c, "implementation/AddOrder",
		trace.WithAttributes(
			attribute.String("Description", "AddOrder in order implementation"),
		),
	)
	defer span.End()

	if err := r.p.ProductDb.Debug().Create(&order).Error; err != nil {
		span.RecordError(err)
		return nil, err
	}

	return order, nil
}

func (r orderRepo) AddOrderTx(tx *gorm.DB, order *entity.Order) (*entity.Order, error) {

	tracer := otel.Tracer("quqo")

	// Start a new span for the function
	_, span := tracer.Start(*r.c, "implementation/AddOrderTx",
		trace.WithAttributes(
			attribute.String("Description", "AddOrderTx in order implementation"),
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

	tracer := otel.Tracer("quqo")
	_, span := tracer.Start(*r.c, "implementation/GetOrder",
		trace.WithAttributes(
			attribute.String("Description", "GetOrder in order implementation"),
		),
	)
	defer span.End()

	var order *entity.Order

	err := r.p.ProductDb.Debug().Unscoped().Preload("OrderedItems").Where("order_id = ?", id).Take(&order).Error

	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		span.RecordError(errors.New("order not found"))
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r orderRepo) UpdateOrder(cus *entity.Order) (*entity.Order, error) {

	tracer := otel.Tracer("quqo")
	_, span := tracer.Start(*r.c, "implementation/UpdateOrder",
		trace.WithAttributes(
			attribute.String("Description", "UpdateOrder in order implementation"),
		),
	)
	defer span.End()

	result := r.p.ProductDb.Debug().Where("order_id = ?", cus.OrderID).Updates(&cus)

	if result.Error != nil {
		span.RecordError(result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		span.RecordError(errors.New("order not found"))
		return nil, errors.New("order not found")
	}

	return cus, nil
}

func (r orderRepo) DeleteOrder(id uint64) error {

	tracer := otel.Tracer("quqo")
	_, span := tracer.Start(*r.c, "implementation/DeleteOrder",
		trace.WithAttributes(
			attribute.String("Description", "DeleteOrder in order implementation"),
		),
	)
	defer span.End()

	var order entity.Order
	res := r.p.ProductDb.Debug().Where("order_id = ?", id).Delete(&order)
	if res.Error != nil {
		span.RecordError(res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		span.RecordError(errors.New("order not found"))
		return errors.New("order not found")
	}

	return nil
}
