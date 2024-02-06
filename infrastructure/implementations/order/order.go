package order

import (
	"errors"
	entity "products-crud/domain/entity/order_entity"
	repository "products-crud/domain/repository/order_repository"

	base "products-crud/infrastructure/persistences"

	"gorm.io/gorm"
)

type orderRepo struct {
	p *base.Persistence
}

func NewOrderRepository(p *base.Persistence) repository.OrderRepository {
	return &orderRepo{p}
}

func (r orderRepo) AddOrder(cus *entity.Order) (*entity.Order, error) {

	if err := r.p.ProductDb.Debug().Create(&cus).Error; err != nil {
		return nil, err
	}

	return cus, nil
}

func (r orderRepo) GetOrder(id uint64) (*entity.Order, error) {
	var order *entity.Order

	err := r.p.ProductDb.Debug().Unscoped().Where("order_id = ?", id).Take(&order).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("order not found")
	}

	return order, nil
}

func (r orderRepo) UpdateOrder(cus *entity.Order) (*entity.Order, error) {
	result := r.p.ProductDb.Debug().Where("order_id = ?", cus.OrderID).Updates(&cus)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("order not found")
	}

	return cus, nil
}

func (r orderRepo) DeleteOrder(id uint64) error {
	var order entity.Order
	res := r.p.ProductDb.Debug().Where("order_id = ?", id).Delete(&order)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}
