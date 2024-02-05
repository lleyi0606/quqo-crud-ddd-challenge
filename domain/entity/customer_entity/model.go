package customerentity

import (
	"products-crud/domain/entity"
)

type Customer struct {
	entity.BaseModelWDelete
	CustomerID uint64  `gorm:"autoIncrement" json:"customer_id"`
	Name       string  `gorm:"size:100" json:"name"`
	Address    string  `gorm:"size:255" json:"address"`
	Longitude  float64 `gorm:"type:double precision" json:"longitude"`
	Latitude   float64 `gorm:"type:double precision" json:"latitude"`
}
