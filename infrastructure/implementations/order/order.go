package order

import (
	"errors"
	entity "products-crud/domain/entity/order_entity"
	repository "products-crud/domain/repository/order_repository"

	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderRepo struct {
	p *base.Persistence
	c *gin.Context
}

func NewOrderRepository(p *base.Persistence, c *gin.Context) repository.OrderRepository {
	return &orderRepo{p, c}
}

func (r orderRepo) AddOrder(order *entity.Order) (*entity.Order, error) {

	span := r.p.Logger.Start(r.c, "infrastructure/implementations/AddOrder", map[string]interface{}{"order": order})
	defer span.End()

	if err := r.p.ProductDb.Debug().Create(&order).Error; err != nil {
		r.p.Logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, err
	}

	return order, nil
}

func (r orderRepo) AddOrderTx(tx *gorm.DB, order *entity.Order) (*entity.Order, error) {

	span := r.p.Logger.Start(r.c, "infrastructure/implementations/AddOrderTx", map[string]interface{}{"order": order})
	defer span.End()

	if err := tx.Debug().Create(&order).Error; err != nil {
		r.p.Logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, err
	}

	return order, nil
}

func (r orderRepo) GetOrder(id uint64) (*entity.Order, error) {

	span := r.p.Logger.Start(r.c, "infrastructure/implementations/GetOrder", map[string]interface{}{"id": id})
	defer span.End()

	var order *entity.Order

	err := r.p.ProductDb.Debug().Unscoped().Preload("OrderedItems").Where("order_id = ?", id).Take(&order).Error

	if err != nil {
		r.p.Logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, err
	}

	r.p.Logger.Info("get order", map[string]interface{}{"data": order})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.p.Logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r orderRepo) UpdateOrder(order *entity.Order) (*entity.Order, error) {

	span := r.p.Logger.Start(r.c, "infrastructure/implementations/UpdateOrder", map[string]interface{}{"order": order})
	defer span.End()

	result := r.p.ProductDb.Debug().Where("order_id = ?", order.OrderID).Updates(&order)

	if result.Error != nil {
		r.p.Logger.Error(result.Error.Error(), map[string]interface{}{"data": order})
		return nil, result.Error
	}

	r.p.Logger.Info("get order", map[string]interface{}{"data": result})

	if result.RowsAffected == 0 {
		r.p.Logger.Error("order not found", map[string]interface{}{"data": order})
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r orderRepo) DeleteOrder(id uint64) error {

	span := r.p.Logger.Start(r.c, "infrastructure/implementations/DeleteOrder", map[string]interface{}{"id": id})
	defer span.End()

	var order entity.Order
	res := r.p.ProductDb.Debug().Where("order_id = ?", id).Delete(&order)
	if res.Error != nil {
		r.p.Logger.Error(res.Error.Error(), map[string]interface{}{"data": order})
		return res.Error
	}

	if res.RowsAffected == 0 {
		r.p.Logger.Error("order not found", map[string]interface{}{"data": order})
		return errors.New("order not found")
	}

	return nil
}
