package base

import (
	"log"
	"products-crud/domain/entity"
	"products-crud/infrastructure/persistences/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type Persistence struct {
	ProductDb *gorm.DB
}

func NewPersistence() (*Persistence, error) {

	// Product engine
	productEngine, errProductE := db.NewProductDB()
	if errProductE != nil {
		log.Fatal(errProductE)
	}

	// Product Redis engine

	return &Persistence{
		ProductDb: productEngine.DB,
	}, nil
}

// closes the  database connection
func (p *Persistence) Close() error {
	pdtDB, errQ := p.ProductDb.DB()
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
	return p.ProductDb.AutoMigrate(&entity.Product{})
}
