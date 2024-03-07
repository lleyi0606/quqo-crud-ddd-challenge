package order

import (
	"errors"
	loggerentity "products-crud/domain/entity/logger_entity"
	entity "products-crud/domain/entity/order_entity"
	repository "products-crud/domain/repository/order_repository"

	"products-crud/infrastructure/implementations/logger"
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

	info := loggerentity.FunctionInfo{
		FunctionName: "AddOrder",
		Path:         "infrastructure/implementations/",
		Description:  "Add order into DB",
	}
	logger, endFunc := logger.NewLoggerRepositories(r.p, r.c, info, []string{"Honeycomb", "zap"})
	defer endFunc()

	if err := r.p.ProductDb.Debug().Create(&order).Error; err != nil {
		logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, err
	}
	logger.End()
	return order, nil
}

func (r orderRepo) AddOrderTx(tx *gorm.DB, order *entity.Order) (*entity.Order, error) {

	info := loggerentity.FunctionInfo{
		FunctionName: "AddOrderTx",
		Path:         "infrastructure/implementations/",
		Description:  "Add order into DB",
		Body:         nil,
	}
	logger, endFunc := logger.NewLoggerRepositories(r.p, r.c, info, []string{"Honeycomb", "zap"})
	defer endFunc()

	if err := tx.Debug().Create(&order).Error; err != nil {
		logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, err
	}

	return order, nil
}

func (r orderRepo) GetOrder(id uint64) (*entity.Order, error) {

	info := loggerentity.FunctionInfo{
		FunctionName: "GetOrder",
		Path:         "infrastructure/implementations/",
		Description:  "Get order from DB",
	}
	logger, endFunc := logger.NewLoggerRepositories(r.p, r.c, info, []string{"Honeycomb", "zap"})
	defer endFunc()

	var order *entity.Order

	err := r.p.ProductDb.Debug().Unscoped().Preload("OrderedItems").Where("order_id = ?", id).Take(&order).Error

	if err != nil {
		logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(err.Error(), map[string]interface{}{"data": order})
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r orderRepo) UpdateOrder(order *entity.Order) (*entity.Order, error) {

	info := loggerentity.FunctionInfo{
		FunctionName: "UpdateOrder",
		Path:         "infrastructure/implementations/",
		Description:  "Update order from DB",
	}
	logger, endFunc := logger.NewLoggerRepositories(r.p, r.c, info, []string{"Honeycomb", "zap"})
	defer endFunc()

	result := r.p.ProductDb.Debug().Where("order_id = ?", order.OrderID).Updates(&order)

	if result.Error != nil {
		logger.Error(result.Error.Error(), map[string]interface{}{"data": order})
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		logger.Error("order not found", map[string]interface{}{"data": order})
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r orderRepo) DeleteOrder(id uint64) error {

	info := loggerentity.FunctionInfo{
		FunctionName: "DeleteOrder",
		Path:         "infrastructure/implementations/",
		Description:  "Delete order from DB",
	}
	logger, endFunc := logger.NewLoggerRepositories(r.p, r.c, info, []string{"Honeycomb", "zap"})
	defer endFunc()

	var order entity.Order
	res := r.p.ProductDb.Debug().Where("order_id = ?", id).Delete(&order)
	if res.Error != nil {
		logger.Error(res.Error.Error(), map[string]interface{}{"data": order})
		return res.Error
	}

	if res.RowsAffected == 0 {
		logger.Error("order not found", map[string]interface{}{"data": order})
		return errors.New("order not found")
	}

	return nil
}
