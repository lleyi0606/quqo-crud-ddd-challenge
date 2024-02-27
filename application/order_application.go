package application

import (
	"context"
	entity "products-crud/domain/entity/order_entity"
	orderItem_entity "products-crud/domain/entity/orderedItem_entity"

	repository "products-crud/domain/repository/order_repository"
	"products-crud/infrastructure/implementations/inventory"
	"products-crud/infrastructure/implementations/order"
	"products-crud/infrastructure/implementations/orderedItem"
	"products-crud/infrastructure/implementations/product"
	base "products-crud/infrastructure/persistences"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type OrderApp struct {
	p *base.Persistence
	c *context.Context
}

func NewOrderApplication(p *base.Persistence, c *context.Context) repository.OrderHandlerRepository {
	return &OrderApp{p, c}
}

func (u *OrderApp) AddOrder(orderInput *entity.OrderInput) (*entity.Order, error) {

	tracer := otel.Tracer("quqo")
	context, span := tracer.Start(*u.c, "application/AddOrder",
		trace.WithAttributes(
			attribute.String("Description", "AddOrder in order application"),
		),
	)
	defer span.End()

	tx := u.p.ProductDb.Begin()
	var errTx error

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if errTx != nil {
			span.RecordError(errTx)
			tx.Rollback()
		} else {
			errC := tx.Commit().Error
			if errC != nil {
				span.RecordError(errC)
				tx.Rollback()
			}
		}
	}()

	repoOrder := order.NewOrderRepository(u.p, &context)

	// update stock
	repoInventory := inventory.NewInventoryRepository(u.p, &context)
	for _, orderedItemInput := range orderInput.OrderedItems {
		errTx = repoInventory.DecreaseStockTx(tx, orderedItemInput.ProductID, orderedItemInput.Quantity)
		if errTx != nil {
			span.RecordError(errTx)
			return nil, errTx
		}
	}

	// add the orderedItems
	var orderedItems []orderItem_entity.OrderedItem
	cost := 0.0

	repoOrderedItem := orderedItem.NewOrderedItemRepository(u.p, &context)
	repoProduct := product.NewProductRepository(u.p, &context)
	for _, orderedItemInput := range orderInput.OrderedItems {
		unitPrice, totalPrice, errTx := repoProduct.CalculateProductPriceByQuantityTx(tx, orderedItemInput.ProductID, orderedItemInput.Quantity)
		if errTx != nil {
			span.RecordError(errTx)
			return nil, errTx
		}
		orderedItem := &orderItem_entity.OrderedItem{
			OrderID:    orderInput.OrderID,
			ProductID:  orderedItemInput.ProductID,
			Quantity:   orderedItemInput.Quantity,
			UnitPrice:  unitPrice,  // You need to set the appropriate value
			TotalPrice: totalPrice, // You need to set the appropriate value
		}

		if _, errTx = repoOrderedItem.AddOrderedItemTx(tx, orderedItem); errTx != nil {
			span.RecordError(errTx)
			return nil, errTx
		}

		cost += totalPrice
		orderedItems = append(orderedItems, *orderedItem)
	}

	// calculate fees
	fees, _ := u.CalculateFees(cost)

	// create and add order
	order := &entity.Order{
		OrderID:       orderInput.OrderID,
		CustomerID:    orderInput.CustomerID,
		WarehouseID:   orderInput.WarehouseID,
		Status:        "Processing",
		OrderedItems:  orderedItems,
		TotalCost:     cost,
		TotalFees:     fees,
		TotalCheckout: cost + fees,
	}
	res, errTx := repoOrder.AddOrderTx(tx, order)
	if errTx != nil {
		span.RecordError(errTx)
		return nil, errTx
	}

	return res, nil
}

func (u *OrderApp) GetOrder(id uint64) (*entity.Order, error) {

	tracer := otel.Tracer("quqo")
	context, span := tracer.Start(*u.c, "application/GetOrder",
		trace.WithAttributes(
			attribute.String("Description", "GetOrder in order application"),
		),
	)
	defer span.End()

	repoOrder := order.NewOrderRepository(u.p, &context)
	return repoOrder.GetOrder(id)
}

func (u *OrderApp) UpdateOrder(cat *entity.Order) (*entity.Order, error) {

	tracer := otel.Tracer("quqo")
	context, span := tracer.Start(*u.c, "application/UpdateOrder",
		trace.WithAttributes(
			attribute.String("Description", "UpdateOrder in order application"),
		),
	)
	defer span.End()

	repoOrder := order.NewOrderRepository(u.p, &context)
	return repoOrder.UpdateOrder(cat)
}

func (u *OrderApp) DeleteOrder(id uint64) error {
	tracer := otel.Tracer("quqo")
	context, span := tracer.Start(*u.c, "application/DeleteOrder",
		trace.WithAttributes(
			attribute.String("Description", "DeleteOrder in order application"),
		),
	)
	defer span.End()

	repoOrder := order.NewOrderRepository(u.p, &context)
	return repoOrder.DeleteOrder(id)
}

func (u *OrderApp) CalculateFees(amt float64) (float64, error) {
	tracer := otel.Tracer("quqo")
	_, span := tracer.Start(*u.c, "application/CalculateFees",
		trace.WithAttributes(
			attribute.String("Description", "CalculateFees in order application"),
		),
	)
	defer span.End()

	return 0.02 * amt, nil
}
