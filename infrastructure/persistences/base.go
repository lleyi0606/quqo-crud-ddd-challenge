package base

import (
	"log"
	"products-crud/domain/entity"
	"products-crud/infrastructure/persistences/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Persistence struct {
	db *gorm.DB
}

func NewPersistence() (*Persistence, error) {

	// Product engine
	productEngine, errProductE := db.NewProductDB()
	if errProductE != nil {
		log.Fatal(errProductE)
	}

	// Product Redis engine

	return &Persistence{
		db: productEngine.DB,
	}, nil
}

// closes the  database connection
func (p *Persistence) Close() error {
	pdtDB, errQ := p.db.DB()
	if errQ != nil {
		return errQ
	}
	errDbClose := pdtDB.Close()
	if errDbClose != nil {
		return errDbClose
	}

	return nil
}

// This migrate all tables
func (p *Persistence) Automigrate() error {
	return p.db.AutoMigrate(&entity.Product{})
}
