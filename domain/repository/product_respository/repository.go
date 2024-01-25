package repository

import entity "products-crud/domain/entity/product_entity"

type ProductRepository interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetProduct(uint64) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	UpdateProduct(*entity.Product) (*entity.Product, error)
	DeleteProduct(uint64) (*entity.Product, error)
	SearchProducts(string) ([]entity.Product, error)
	// GetProductsByInventory(uint64) ([]entity.Product, error)
}

type ProductHandlerRepository interface {
	AddProduct(*entity.ProductWithStockAndWarehouse) (*entity.Product, error)
	GetProduct(uint64) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	UpdateProduct(*entity.Product) (*entity.Product, error)
	DeleteProduct(uint64) (*entity.Product, error)
	SearchProducts(string) ([]entity.Product, error)
}
