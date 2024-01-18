package product

import (
	"errors"
	"log"
	"products-crud/domain/entity"
	base "products-crud/infrastructure/persistences"

	"gorm.io/gorm"
)

type productRepo struct {
	p *base.Persistence
}

func NewProductRepository(p *base.Persistence) *productRepo {
	return &productRepo{p}
}

// productRepo implements the repository.ProductRepository interface

func (r productRepo) AddProduct(pdt *entity.Product) (*entity.Product, error) {
	log.Println("Adding new product ", pdt.Name, "...")

	if err := r.p.ProductDb.Debug().Create(&pdt).Error; err != nil {
		return nil, err
	}

	log.Println(pdt.Name, " created.")
	return pdt, nil
}

func (r productRepo) GetProduct(id uint64) (*entity.Product, error) {
	var pdt entity.Product
	err := r.p.ProductDb.Debug().Where("id = ?", id).Take(&pdt).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return &pdt, nil
}

func (r productRepo) GetProducts() ([]entity.Product, error) {
	var pdts []entity.Product
	err := r.p.ProductDb.Debug().Find(&pdts).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return pdts, nil
}

func (r productRepo) UpdateProduct(pdt *entity.ProductUpdate) (*entity.Product, error) {
	err := r.p.ProductDb.Debug().Model(&entity.Product{}).Where("id = ?", pdt.ID).Updates(pdt).Error

	if err != nil {
		return nil, err
	}

	return r.GetProduct(pdt.ID)
}

func (r productRepo) DeleteProduct(id uint64) (*entity.Product, error) {
	var pdt entity.Product
	err := r.p.ProductDb.Debug().Where("id = ?", id).Delete(&pdt).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return &pdt, nil
}

func (r productRepo) SearchProducts(str string) ([]entity.Product, error) {
	var pdts []entity.Product
	err := r.p.ProductDb.Debug().Where("lower(name) LIKE lower(?)", "%"+str+"%").Find(&pdts).Error
	if err != nil {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return pdts, nil
}
