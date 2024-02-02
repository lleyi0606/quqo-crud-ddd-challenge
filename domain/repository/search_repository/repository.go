package search_repository

import (
	entity "products-crud/domain/entity/product_entity"
)

type SearchRepository interface {
	AddProduct(*entity.Product) error
	SearchProducts(string) ([]entity.Product, error)
	DeleteProduct(uint64) error
	// UpdateProduct(*entity.Product) error
}
