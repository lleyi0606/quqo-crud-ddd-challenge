package categoryentity

import (
	"products-crud/domain/entity"
)

type Category struct {
	entity.BaseModelWDelete
	CategoryID uint64 `gorm:"primaryKey;autoIncrement:true" json:"category_id"`
	Name       string `gorm:"size:255;" json:"name"`
	ParentID   int64  `gorm:"type:numeric;" json:"parent_id"`
}
