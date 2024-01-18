package application

import (
	"products-crud/domain/entity"
	"products-crud/domain/repository"
)

type productApp struct {
	us repository.ProductRepository
}

// ProductApp implements the ProductAppInterface
var _ ProductAppInterface = &productApp{}

type ProductAppInterface interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetProduct(uint64) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	UpdateProduct(*entity.ProductUpdate) (*entity.Product, error)
	DeleteProduct(uint64) (*entity.Product, error)
	SearchProducts(string) ([]entity.Product, error)
}

func (u *productApp) AddProduct(user *entity.Product) (*entity.Product, error) {
	return u.us.AddProduct(user)
}

func (u *productApp) GetProduct(pdtId uint64) (*entity.Product, error) {
	return u.us.GetProduct(pdtId)
}

func (u *productApp) GetProducts() ([]entity.Product, error) {
	return u.us.GetProducts()
}

func (u *productApp) UpdateProduct(pdt *entity.ProductUpdate) (*entity.Product, error) {
	return u.us.UpdateProduct(pdt)
}

func (u *productApp) DeleteProduct(pdtId uint64) (*entity.Product, error) {
	return u.us.DeleteProduct(pdtId)
}

func (u *productApp) SearchProducts(str string) ([]entity.Product, error) {
	return u.us.SearchProducts(str)
}
