package orderedItem

import (
	"errors"
	entity "products-crud/domain/entity/orderedItem_entity"
	repository "products-crud/domain/repository/orderedItem_repository"

	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type orderedItemRepo struct {
	p *base.Persistence
	c *gin.Context
}

func NewOrderedItemRepository(p *base.Persistence, c *gin.Context) repository.OrderedItemRepository {
	return &orderedItemRepo{p, c}
}

func (r orderedItemRepo) AddOrderedItem(item *entity.OrderedItem) (*entity.OrderedItem, error) {

	if err := r.p.ProductDb.Debug().Create(&item).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (r orderedItemRepo) AddOrderedItemTx(tx *gorm.DB, item *entity.OrderedItem) (*entity.OrderedItem, error) {

	// info := loggerentity.FunctionInfo{
	// 	FunctionName: "AddOrderedItemTx",
	// 	Path:         "infrastructure/implementations/",
	// 	Description:  "Add ordered item into DB",
	// 	Body:         nil,
	// }
	// logger, endFunc := logger.NewLoggerRepositories(r.p, r.c, info, []string{"Honeycomb", "zap"})
	// defer endFunc()

	if err := tx.Debug().Create(&item).Error; err != nil {
		// logger.Error(err.Error(), map[string]interface{}{})
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
