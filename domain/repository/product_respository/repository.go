package repository

import (
	entity "products-crud/domain/entity/product_entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetProduct(string) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	UpdateProduct(*entity.Product) (*entity.Product, error)
	DeleteProduct(string) (*entity.Product, error)
	// SearchProducts(string) ([]entity.Product, error)
	CalculateProductPriceByQuantity(string, int) (float64, float64, error)
	CalculateProductPriceByQuantityTx(*gorm.DB, string, int) (float64, float64, error)
	// GetProductsByInventory(string) ([]entity.Product, error)
}

type ProductHandlerRepository interface {
	AddProduct(*entity.ProductWithStockAndWarehouse) (*entity.Product, error)
	GetProduct(string) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	UpdateProduct(*entity.Product) (*entity.Product, error)
	DeleteProduct(string) (*entity.Product, error)
	SearchProducts(string) ([]entity.Product, error)
}
