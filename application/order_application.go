package application

import (
	entity "products-crud/domain/entity/order_entity"
	orderItem_entity "products-crud/domain/entity/orderedItem_entity"

	repository "products-crud/domain/repository/order_repository"
	"products-crud/infrastructure/implementations/inventory"
	"products-crud/infrastructure/implementations/order"
	"products-crud/infrastructure/implementations/orderedItem"
	"products-crud/infrastructure/implementations/product"
	base "products-crud/infrastructure/persistences"
)

type OrderApp struct {
	p *base.Persistence
}

func NewOrderApplication(p *base.Persistence) repository.OrderHandlerRepository {
	return &OrderApp{p}
}

func (u *OrderApp) AddOrder(orderInput *entity.OrderInput) (*entity.Order, error) {

	tx := u.p.ProductDb.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	repoOrder := order.NewOrderRepository(u.p)

	// update stock
	repoInventory := inventory.NewInventoryRepository(u.p)
	for _, orderedItemInput := range orderInput.OrderedItems {
		err := repoInventory.DecreaseStockTx(tx, orderedItemInput.ProductID, orderedItemInput.Quantity)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// add the orderedItems
	var orderedItems []orderItem_entity.OrderedItem
	cost := 0.0

	repoOrderedItem := orderedItem.NewOrderedItemRepository(u.p)
	repoProduct := product.NewProductRepository(u.p)
	for _, orderedItemInput := range orderInput.OrderedItems {
		unitPrice, totalPrice, err := repoProduct.CalculateProductPriceByQuantityTx(tx, orderedItemInput.ProductID, orderedItemInput.Quantity)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		orderedItem := &orderItem_entity.OrderedItem{
			OrderID:    orderInput.OrderID,
			ProductID:  orderedItemInput.ProductID,
			Quantity:   orderedItemInput.Quantity,
			UnitPrice:  unitPrice,  // You need to set the appropriate value
			TotalPrice: totalPrice, // You need to set the appropriate value
		}

		if _, err := repoOrderedItem.AddOrderedItemTx(tx, orderedItem); err != nil {
			tx.Rollback()
			return nil, err
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
	res, err := repoOrder.AddOrderTx(tx, order)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return res, tx.Commit().Error
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

func (u *OrderApp) CalculateFees(amt float64) (float64, error) {
	return 0.02 * amt, nil
}
