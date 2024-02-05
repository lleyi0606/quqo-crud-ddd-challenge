package category

import (
	"errors"
	"log"
	entity "products-crud/domain/entity/category_entity"
	repository "products-crud/domain/repository/category_repository"

	base "products-crud/infrastructure/persistences"

	"gorm.io/gorm"
)

type categoryRepo struct {
	p *base.Persistence
}

func NewCategoryRepository(p *base.Persistence) repository.CategoryRepository {
	return &categoryRepo{p}
}

// categoryRepo implements the repository.categoryRepository interface

func (r categoryRepo) AddCategory(cat *entity.Category) (*entity.Category, error) {
	log.Println("Adding new category ", cat.Name, "...")

	if err := r.p.ProductDb.Debug().Create(&cat).Error; err != nil {
		return nil, err
	}

	log.Println(cat.Name, " created.")
	return cat, nil
}

func (r categoryRepo) GetCategory(id uint64) (*entity.Category, error) {
	var cat *entity.Category

	err := r.p.ProductDb.Debug().Unscoped().Where("category_id = ?", id).Take(&cat).Error
	if err != nil {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("category not found")
	}

	return cat, nil
}

// func (r categoryRepo) GetCategories() ([]entity.Category, error) {
// 	var cats []entity.Category
// 	err := r.p.ProductDb.Debug().Find(&cats).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, errors.New("category not found")
// 	}
// 	return cats, nil
// }

func (r categoryRepo) UpdateCategory(cat *entity.Category) (*entity.Category, error) {
	// err := r.p.categoryDb.Debug().Model(&entity.category{}).Where("id = ?", cat.ID).Updates(cat).Error
	err := r.p.ProductDb.Debug().Where("category_id = ?", cat.CategoryID).Updates(&cat).Error

	if err != nil {
		return nil, err
	}

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return cat, nil
}

func (r categoryRepo) DeleteCategory(id uint64) error {
	var cat entity.Category
	res := r.p.ProductDb.Debug().Where("category_id = ?", id).Delete(&cat)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}
