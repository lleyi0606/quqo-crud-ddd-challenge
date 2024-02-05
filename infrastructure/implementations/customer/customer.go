package customer

import (
	"errors"
	"log"
	entity "products-crud/domain/entity/customer_entity"
	repository "products-crud/domain/repository/customer_repository"

	base "products-crud/infrastructure/persistences"

	"gorm.io/gorm"
)

type customerRepo struct {
	p *base.Persistence
}

func NewCustomerRepository(p *base.Persistence) repository.CustomerRepository {
	return &customerRepo{p}
}

func (r customerRepo) AddCustomer(cus *entity.Customer) (*entity.Customer, error) {
	log.Println("Adding new customer ", cus.Name, "...")

	if err := r.p.ProductDb.Debug().Create(&cus).Error; err != nil {
		return nil, err
	}

	log.Println(cus.Name, " created.")
	return cus, nil
}

func (r customerRepo) GetCustomer(id uint64) (*entity.Customer, error) {
	var cus *entity.Customer

	err := r.p.ProductDb.Debug().Unscoped().Where("customer_id = ?", id).Take(&cus).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("customer not found")
	}

	return cus, nil
}

func (r customerRepo) UpdateCustomer(cus *entity.Customer) (*entity.Customer, error) {
	err := r.p.ProductDb.Debug().Where("customer_id = ?", cus.CustomerID).Updates(&cus).Error

	if err != nil {
		return nil, err
	}

	return cus, nil
}

func (r customerRepo) DeleteCustomer(id uint64) error {
	var cus entity.Customer
	res := r.p.ProductDb.Debug().Where("customer_id = ?", id).Delete(&cus)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("customer not found")
	}

	return nil
}
