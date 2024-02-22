package repository

import (
	entity "products-crud/domain/entity/customer_entity"
)

type CustomerRepository interface {
	AddCustomer(*entity.Customer) (*entity.Customer, error)
	GetCustomer(uint64) (*entity.Customer, error)
	UpdateCustomer(*entity.Customer) (*entity.Customer, error)
	DeleteCustomer(uint64) error
	GetCustomerByUsernameAndPassword(*entity.Customer) (*entity.Customer, error)
}

type CustomerHandlerRepository interface {
	AddCustomer(*entity.Customer) (*entity.Customer, error)
	GetCustomer(uint64) (*entity.Customer, error)
	UpdateCustomer(*entity.Customer) (*entity.Customer, error)
	DeleteCustomer(uint64) error
}
