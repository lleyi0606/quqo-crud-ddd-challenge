package application

import (
	loggerentity "products-crud/domain/entity/logger_entity"
	entity "products-crud/domain/entity/order_entity"
	orderItem_entity "products-crud/domain/entity/orderedItem_entity"

	repository "products-crud/domain/repository/order_repository"
	"products-crud/infrastructure/implementations/inventory"
	"products-crud/infrastructure/implementations/logger"
	"products-crud/infrastructure/implementations/order"
	"products-crud/infrastructure/implementations/orderedItem"
	"products-crud/infrastructure/implementations/product"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

type OrderApp struct {
	p *base.Persistence
	c *gin.Context
}

func NewOrderApplication(p *base.Persistence, c *gin.Context) repository.OrderHandlerRepository {
	return &OrderApp{p, c}
}

func (u *OrderApp) AddOrder(orderInput *entity.OrderInput) (*entity.Order, error) {

	info := loggerentity.FunctionInfo{
		FunctionName: "AddOrder",
		Path:         "application/",
		Description:  "Application of add order",
	}
	logger := logger.NewLoggerRepositories(u.p, u.c, info, "honeycomb", "zap")
	defer logger.End()

	tx := u.p.ProductDb.Begin()
	var errTx error

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if errTx != nil {
			tx.Rollback()
		} else {
			errC := tx.Commit().Error
			if errC != nil {
				tx.Rollback()
			}
		}
	}()

	// update stock
	repoInventory := inventory.NewInventoryRepository(u.p, u.c)
	for _, orderedItemInput := range orderInput.OrderedItems {
		errTx = repoInventory.DecreaseStockTx(tx, orderedItemInput.ProductID, orderedItemInput.Quantity)
		if errTx != nil {
			return nil, errTx
		}
	}

	// add the orderedItems
	var orderedItems []orderItem_entity.OrderedItem
	cost := 0.0

	repoOrderedItem := orderedItem.NewOrderedItemRepository(u.p, u.c)
	repoProduct := product.NewProductRepository(u.p, u.c)
	for _, orderedItemInput := range orderInput.OrderedItems {
		unitPrice, totalPrice, errTx := repoProduct.CalculateProductPriceByQuantityTx(tx, orderedItemInput.ProductID, orderedItemInput.Quantity)
		if errTx != nil {
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
			return nil, errTx
		}

		cost += totalPrice
		orderedItems = append(orderedItems, *orderedItem)
	}

	// calculate fees
	fees, _ := u.CalculateFees(cost)

	// create and add order
	newOrder := &entity.Order{
		OrderID:       orderInput.OrderID,
		CustomerID:    orderInput.CustomerID,
		WarehouseID:   orderInput.WarehouseID,
		Status:        "Processing",
		OrderedItems:  orderedItems,
		TotalCost:     cost,
		TotalFees:     fees,
		TotalCheckout: cost + fees,
	}

	repoOrder := order.NewOrderRepository(u.p, u.c)
	res, errTx := repoOrder.AddOrderTx(tx, newOrder)
	if errTx != nil {
		return nil, errTx
	}

	return res, nil
}

func (u *OrderApp) GetOrder(id uint64) (*entity.Order, error) {

	repoOrder := order.NewOrderRepository(u.p, u.c)
	return repoOrder.GetOrder(id)
}

func (u *OrderApp) UpdateOrder(cat *entity.Order) (*entity.Order, error) {

	repoOrder := order.NewOrderRepository(u.p, u.c)
	return repoOrder.UpdateOrder(cat)
}

func (u *OrderApp) DeleteOrder(id uint64) error {

	repoOrder := order.NewOrderRepository(u.p, u.c)
	return repoOrder.DeleteOrder(id)
}

func (u *OrderApp) CalculateFees(amt float64) (float64, error) {

	return 0.02 * amt, nil
}
