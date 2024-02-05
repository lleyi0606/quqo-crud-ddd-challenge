package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type BaseModelWOutDates struct {
	ID int64 `json:"id" gorm:"primaryKey"`
}

type BaseModelWOutID struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type BaseModelWDelete struct {
	ID        int64          `json:"id"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" db:"deleted_at"`
}
