package customerentity

import (
	"products-crud/domain/entity"
	"products-crud/infrastructure/utils"

	"gorm.io/gorm"
)

type Customer struct {
	Username string `gorm:"unique" json:"username"`
	Password string `gorm:"size:100" json:"password"`
	PublicCustomer
}

type PublicCustomer struct {
	CustomerID uint64  `gorm:"autoIncrement" json:"customer_id"`
	Name       string  `gorm:"size:100" json:"name"`
	Address    string  `gorm:"size:255" json:"address"`
	Longitude  float64 `gorm:"type:double precision" json:"longitude"`
	Latitude   float64 `gorm:"type:double precision" json:"latitude"`
	entity.BaseModelWDelete
}

// BeforeSave is a gorm hook
func (u *Customer) BeforeSave(tx *gorm.DB) error {
	hashPassword, err := utils.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}
