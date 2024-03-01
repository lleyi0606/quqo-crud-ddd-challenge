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

	// tracer := otel.Tracer("quqo")
	// _, span := tracer.Start(*r.c, "implementation/AddOrder",
	// 	trace.WithAttributes(
	// 		attribute.String("Description", "AddOrder in order implementation"),
	// 	),
	// )
	// defer span.End()

	// info := loggerentity.FunctionInfo{
	// 	FunctionName: "AddOrder",
	// 	Path:         "infrastructure/implementations/",
	// 	Description:  "Adds order to SQL database",
	// 	Body:         order,
	// }
	// logger := logger.NewLoggerRepositories(r.p, r.c, info, "honeycomb", "zap")
	// defer logger.End()

	if err := r.p.ProductDb.Debug().Create(&order).Error; err != nil {
		// logger.LogError(span, err)
		return nil, err
	}

	return order, nil
}

func (r orderRepo) AddOrderTx(tx *gorm.DB, order *entity.Order) (*entity.Order, error) {

	info := loggerentity.FunctionInfo{
		FunctionName: "AddOrderTx",
		Path:         "infrastructure/implementations/",
		Description:  "Add order into DB",
		Body:         nil,
	}
	logger := logger.NewLoggerRepositories(r.p, r.c, info, "honeycomb", "zap")
	defer logger.End()

	if err := tx.Debug().Create(&order).Error; err != nil {
		// logger.LogError(span, err)
		return nil, err
	}

	return order, nil
}

func (r orderRepo) GetOrder(id uint64) (*entity.Order, error) {

	// tracer := otel.Tracer("quqo")
	// _, span := tracer.Start(*r.c, "implementation/GetOrder",
	// 	trace.WithAttributes(
	// 		attribute.String("Description", "GetOrder in order implementation"),
	// 	),
	// )
	// defer span.End()

	// logger := logger.NewLoggerRepository(r.p, r.c, "Honeycomb")
	// newSpan := loggerentity.Span{
	// 	FunctionName: "GetOrder",
	// 	Path:         "infrastructure/implementations/order/",
	// 	Description:  "GetOrder in implementation",
	// }
	// _, span := logger.NewSpan(&newSpan)
	// defer logger.EndSpan(span)

	var order *entity.Order

	err := r.p.ProductDb.Debug().Unscoped().Preload("OrderedItems").Where("order_id = ?", id).Take(&order).Error

	if err != nil {
		// logger.LogError(span, err)
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// logger.LogError(span, errors.New("order not found"))
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r orderRepo) UpdateOrder(cus *entity.Order) (*entity.Order, error) {

	// tracer := otel.Tracer("quqo")
	// _, span := tracer.Start(*r.c, "implementation/UpdateOrder",
	// 	trace.WithAttributes(
	// 		attribute.String("Description", "UpdateOrder in order implementation"),
	// 	),
	// )
	// defer span.End()

	// logger := logger.NewLoggerRepository(r.p, r.c, "Honeycomb")
	// newSpan := loggerentity.Span{
	// 	FunctionName: "UpdateOrder",
	// 	Path:         "infrastructure/implementations/order/",
	// 	Description:  "UpdateOrder in implementation",
	// }
	// _, span := logger.NewSpan(&newSpan)
	// defer logger.EndSpan(span)

	result := r.p.ProductDb.Debug().Where("order_id = ?", cus.OrderID).Updates(&cus)

	if result.Error != nil {
		// logger.LogError(span, result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		// logger.LogError(span, errors.New("order not found"))
		return nil, errors.New("order not found")
	}

	return cus, nil
}

func (r orderRepo) DeleteOrder(id uint64) error {

	// tracer := otel.Tracer("quqo")
	// _, span := tracer.Start(*r.c, "implementation/DeleteOrder",
	// 	trace.WithAttributes(
	// 		attribute.String("Description", "DeleteOrder in order implementation"),
	// 	),
	// )
	// defer span.End()

	// logger := logger.NewLoggerRepository(r.p, r.c, "Honeycomb")
	// newSpan := loggerentity.Span{
	// 	FunctionName: "DeleteOrder",
	// 	Path:         "infrastructure/implementations/order/",
	// 	Description:  "DeleteOrder in implementation",
	// }
	// _, span := logger.NewSpan(&newSpan)
	// defer logger.EndSpan(span)

	var order entity.Order
	res := r.p.ProductDb.Debug().Where("order_id = ?", id).Delete(&order)
	if res.Error != nil {
		// logger.LogError(span, res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		// logger.LogError(span, errors.New("order not found"))
		return errors.New("order not found")
	}

	return nil
}
