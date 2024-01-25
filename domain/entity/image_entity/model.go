package entity

type Image struct {
	ImageID   uint64  `gorm:"primaryKey;autoIncrement:true" json:"image_id"`
	ProductID uint64  `gorm:"type:numeric" json:"product_id"`
	Url       int     `gorm:"type:numeric;" json:"stock"`
	Caption   *string `gorm:"size:255;" json:"caption"`
}
