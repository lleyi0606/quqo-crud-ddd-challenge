package repository

import (
	"products-crud/domain/entity"
)

type ProductRepository interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetProduct(uint64) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	UpdateProduct(*entity.ProductUpdate) (*entity.Product, error)
	DeleteProduct(uint64) (*entity.Product, error)
	SearchProducts(string) ([]entity.Product, error)
}

type ProductHandlerRepository interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetProduct(uint64) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	UpdateProduct(*entity.ProductUpdate) (*entity.Product, error)
	DeleteProduct(uint64) (*entity.Product, error)
	SearchProducts(string) ([]entity.Product, error)
}
