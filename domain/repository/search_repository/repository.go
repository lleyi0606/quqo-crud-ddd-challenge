package search_repository

import (
	entity "products-crud/domain/entity/product_entity"
)

type SearchRepository interface {
	AddProduct(*entity.Product) error
	SearchData(string) ([]map[string]interface{}, error)
	DeleteProduct(string) error
	// UpdateProduct(*entity.Product) error
}
