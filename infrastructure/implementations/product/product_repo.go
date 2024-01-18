package product

import (
	"errors"
	"log"
	"products-crud/domain/entity"
	"products-crud/domain/repository"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db}
}

// ProductRepo implements the repository.ProductRepository interface
var _ repository.ProductRepository = &ProductRepo{}

func (r *ProductRepo) AddProduct(pdt *entity.Product) (*entity.Product, error) {
	log.Println("Adding new product ", pdt.Name, "...")

	if err := r.db.Debug().Create(&pdt).Error; err != nil {
		return nil, err
	}

	log.Println(pdt.Name, " created.")
	return pdt, nil
}

func (r *ProductRepo) GetProduct(id uint64) (*entity.Product, error) {
	var pdt entity.Product
	err := r.db.Debug().Where("id = ?", id).Take(&pdt).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return &pdt, nil
}

func (r *ProductRepo) GetProducts() ([]entity.Product, error) {
	var pdts []entity.Product
	err := r.db.Debug().Find(&pdts).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return pdts, nil
}

func (r *ProductRepo) UpdateProduct(pdt *entity.ProductUpdate) (*entity.Product, error) {
	err := r.db.Debug().Model(&entity.Product{}).Where("id = ?", pdt.ID).Updates(pdt).Error

	if err != nil {
		return nil, err
	}

	return r.GetProduct(pdt.ID)
}

func (r *ProductRepo) DeleteProduct(id uint64) (*entity.Product, error) {
	var pdt entity.Product
	err := r.db.Debug().Where("id = ?", id).Delete(&pdt).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return &pdt, nil
}

func (r *ProductRepo) SearchProducts(str string) ([]entity.Product, error) {
	var pdts []entity.Product
	err := r.db.Debug().Where("lower(name) LIKE lower(?)", "%"+str+"%").Find(&pdts).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return pdts, nil
}
