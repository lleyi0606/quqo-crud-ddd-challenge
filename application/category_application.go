package application

import (
	entity "products-crud/domain/entity/category_entity"
	repository "products-crud/domain/repository/category_repository"
	"products-crud/infrastructure/implementations/category"
	base "products-crud/infrastructure/persistences"
)

type categoryApp struct {
	p *base.Persistence
}

func NewCategoryApplication(p *base.Persistence) repository.CategoryHandlerRepository {
	return &categoryApp{p}
}

func (u *categoryApp) AddCategory(cat *entity.Category) (*entity.Category, error) {
	repoCategory := category.NewCategoryRepository(u.p)
	return repoCategory.AddCategory(cat)
}

func (u *categoryApp) GetCategory(id uint64) (*entity.Category, error) {
	repoCategory := category.NewCategoryRepository(u.p)
	return repoCategory.GetCategory(id)
}

func (u *categoryApp) UpdateCategory(cat *entity.Category) (*entity.Category, error) {
	repoCategory := category.NewCategoryRepository(u.p)
	return repoCategory.UpdateCategory(cat)
}

func (u *categoryApp) DeleteCategory(id uint64) error {
	repoCategory := category.NewCategoryRepository(u.p)
	return repoCategory.DeleteCategory(id)
}

func (u *categoryApp) GetCategoryChain(id uint64) ([]entity.Category, error) {
	var chain []entity.Category
	repoCategory := category.NewCategoryRepository(u.p)

	for id != 0 {
		temp, err := repoCategory.GetCategory(id)
		if err != nil {
			return nil, err
		}
		chain = append(chain, *temp)
		id = uint64(temp.ParentID)
	}

	return chain, nil

}
