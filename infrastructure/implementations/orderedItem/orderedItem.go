package orderedItem

import (
	"errors"
	entity "products-crud/domain/entity/orderedItem_entity"
	repository "products-crud/domain/repository/orderedItem_repository"

	base "products-crud/infrastructure/persistences"

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

func (r orderedItemRepo) GetOrderedItems() ([]entity.OrderedItem, error) {
	var orderedItem []entity.OrderedItem

	err := r.p.ProductDb.Debug().Take(&orderedItem).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("orderedItem not found")
	}

	return orderedItem, nil
}

func (r orderedItemRepo) GetOrderedItem(id uint64) (*entity.OrderedItem, error) {
	var orderedItem *entity.OrderedItem

	err := r.p.ProductDb.Debug().Unscoped().Where("orderedItem_id = ?", id).Take(&orderedItem).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("orderedItem not found")
	}

	return orderedItem, nil
}

func (r orderedItemRepo) UpdateOrderedItem(cus *entity.OrderedItem) (*entity.OrderedItem, error) {
	result := r.p.ProductDb.Debug().Where("order_id = ? AND product_id = ?", cus.OrderID, cus.ProductID).Updates(&cus)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("orderedItem not found")
	}

	return cus, nil
}

func (r orderedItemRepo) DeleteOrderedItem(id uint64) error {
	var orderedItem entity.OrderedItem
	res := r.p.ProductDb.Debug().Where("orderedItem_id = ?", id).Delete(&orderedItem)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("orderedItem not found")
	}

	return nil
}
