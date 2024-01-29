package repository

import (
	entity "products-crud/domain/entity/image_entity"
)

type ImageRepository interface {
	AddImage(*entity.ImageInput) (*entity.Image, error)
	GetImage(uint64) ([]entity.Image, error)
	DeleteImage(uint64) error
}

type ImageHandlerRepository interface {
	AddImage(*entity.ImageInput) (*entity.Image, error)
	GetImage(uint64) ([]entity.Image, error)
	DeleteImage(uint64) error
}
