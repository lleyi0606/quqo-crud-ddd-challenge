package persistence

import (
	"log"
	"products-crud/domain/entity"
	"products-crud/domain/repository"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	Product repository.ProductRepository
	db      *gorm.DB
}

func NewRepositories(DbPassword string) (*Repositories, error) {
	dsn := "postgresql://user-test:" + DbPassword + "@fitful-condor-8198.8nk.gcp-asia-southeast1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
		return nil, err
	}

	// DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	// db, err := gorm.Open(Dbdriver, DBURL)
	// if err != nil {
	// 	return nil, err
	// }
	// db.LogMode(true)

	return &Repositories{
		Product: NewProductRepository(db),
		db:      db,
	}, nil
}

// closes the  database connection
// func (s *Repositories) Close() error {
// 	return s.db.Close()
// }

// This migrate all tables
func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.Product{})
}
