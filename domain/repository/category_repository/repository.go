package repository

import (
	entity "products-crud/domain/entity/category_entity"
)

type CategoryRepository interface {
	AddCategory(*entity.Category) (*entity.Category, error)
	GetCategory(uint64) (*entity.Category, error)
	UpdateCategory(*entity.Category) (*entity.Category, error)
	DeleteCategory(uint64) error
}

type CategoryHandlerRepository interface {
	AddCategory(*entity.Category) (*entity.Category, error)
	GetCategory(uint64) (*entity.Category, error)
	GetCategoryChain(uint64) ([]entity.Category, error)
	UpdateCategory(*entity.Category) (*entity.Category, error)
	DeleteCategory(uint64) error
}
