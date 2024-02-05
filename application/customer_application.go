package application

import (
	entity "products-crud/domain/entity/customer_entity"
	repository "products-crud/domain/repository/customer_repository"
	"products-crud/infrastructure/implementations/customer"
	base "products-crud/infrastructure/persistences"
)

type CustomerApp struct {
	p *base.Persistence
}

func NewCustomerApplication(p *base.Persistence) repository.CustomerHandlerRepository {
	return &CustomerApp{p}
}

func (u *CustomerApp) AddCustomer(cat *entity.Customer) (*entity.Customer, error) {
	repoCustomer := customer.NewCustomerRepository(u.p)
	return repoCustomer.AddCustomer(cat)
}

func (u *CustomerApp) GetCustomer(id uint64) (*entity.Customer, error) {
	repoCustomer := customer.NewCustomerRepository(u.p)
	return repoCustomer.GetCustomer(id)
}

func (u *CustomerApp) UpdateCustomer(cat *entity.Customer) (*entity.Customer, error) {
	repoCustomer := customer.NewCustomerRepository(u.p)
	return repoCustomer.UpdateCustomer(cat)
}

func (u *CustomerApp) DeleteCustomer(id uint64) error {
	repoCustomer := customer.NewCustomerRepository(u.p)
	return repoCustomer.DeleteCustomer(id)
}
