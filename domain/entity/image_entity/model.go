package entity

import "mime/multipart"

type Image struct {
	ImageID   uint64 `gorm:"primaryKey;autoIncrement:true" json:"image_id"`
	ProductID uint64 `gorm:"type:numeric" json:"product_id"`
	Url       string `gorm:"size:255;" json:"url"`
	Caption   string `gorm:"size:255;" json:"caption"`
}

type ImageInput struct {
	ProductID uint64                `gorm:"type:numeric" json:"product_id"`
	ImageFile *multipart.FileHeader `form:"image_file" json:"image_file"`
	Caption   string                `form:"caption" gorm:"size:255;" json:"caption"`
}
