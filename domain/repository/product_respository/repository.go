package repository

import (
	"context"
	entity "products-crud/domain/entity/product_entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(*entity.Product) (*entity.Product, error)
	GetProduct(uint64) (*entity.Product, error)
	GetProducts() ([]entity.Product, error)
	UpdateProduct(*entity.Product) (*entity.Product, error)
	DeleteProduct(uint64) (*entity.Product, error)
	SearchProducts(string) ([]entity.Product, error)
	CalculateProductPriceByQuantity(uint64, int) (float64, float64, error)
	CalculateProductPriceByQuantityTx(*gorm.DB, uint64, int, context.Context) (float64, float64, error)
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
