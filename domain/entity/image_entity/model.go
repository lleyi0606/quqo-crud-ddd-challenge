package entity

import (
	"mime/multipart"
	"products-crud/domain/entity"
)

type Image struct {
	entity.BaseModelWDelete
	ImageID   uint64 `gorm:"primaryKey;autoIncrement:true" json:"image_id"`
	ProductID uint64 `gorm:"type:numeric" json:"product_id"`
	Url       string `gorm:"size:255;" json:"url"`
	Caption   string `gorm:"size:255;" json:"caption"`
}

type ImageInput struct {
	entity.BaseModelWDelete
	ProductID uint64                `gorm:"type:numeric" json:"product_id"`
	ImageFile *multipart.FileHeader `form:"image_file" json:"image_file"`
	Caption   string                `form:"caption" gorm:"size:255;" json:"caption"`
}
