package repository

import (
	entity "products-crud/domain/entity/image_entity"
)

type ImageRepository interface {
	AddImage(*entity.Image) (*entity.Image, error)
	// GetInventory(uint64) (*entity.Image, error)
	// GetInventories() ([]entity.Image, error)
	// UpdateInventory(*entity.Image) (*entity.Image, error)
	// DeleteInventory(uint64) (*entity.Image, error)
	// SearchInventory(string) ([]entity.Image, error)
}

type ImageHandlerRepository interface {
	AddImage(*entity.Image) (*entity.Image, error)
	// GetInventory(uint64) (*entity.Image, error)
	// GetInventories() ([]entity.Image, error)
	// UpdateInventory(*entity.Image) (*entity.Image, error)
	// DeleteInventory(uint64) (*entity.Image, error)
	// SearchInventory(string) ([]entity.Image, error)

	// GetInventory(uint64) (*entity.Image, error)
}
