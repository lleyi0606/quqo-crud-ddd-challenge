package orderedItem

import (
	"context"
	"errors"
	entity "products-crud/domain/entity/orderedItem_entity"
	repository "products-crud/domain/repository/orderedItem_repository"

	base "products-crud/infrastructure/persistences"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type orderedItemRepo struct {
	p *base.Persistence
}

func NewOrderedItemRepository(p *base.Persistence) repository.OrderedItemRepository {
	return &orderedItemRepo{p}
}

func (r orderedItemRepo) AddOrderedItem(item *entity.OrderedItem) (*entity.OrderedItem, error) {

	if err := r.p.ProductDb.Debug().Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (r orderedItemRepo) AddOrderedItemTx(tx *gorm.DB, item *entity.OrderedItem, ctx context.Context) (*entity.OrderedItem, error) {

	tracer := otel.Tracer("quqo")

	// Start a new span for the function
	_, span := tracer.Start(ctx, "implementation/AddOrderedItemTx",
		trace.WithAttributes(
			attribute.String("Description", "AddOrderedItemTx in implementation"),
		),
	)
	defer span.End()

	if err := tx.Debug().Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (r orderedItemRepo) GetOrderedItems() ([]entity.OrderedItem, error) {
	var orderedItem []entity.OrderedItem

	err := r.p.ProductDb.Debug().Find(&orderedItem).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("orderedItem not found")
	}

	return orderedItem, nil
}

func (r orderedItemRepo) GetOrderedItemsByOrderId(id uint64) ([]entity.OrderedItem, error) {
	var orderedItems []entity.OrderedItem

	err := r.p.ProductDb.Debug().Unscoped().Where("order_id = ?", id).Take(&orderedItems).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("orderedItems not found")
	}

	return orderedItems, nil
}
