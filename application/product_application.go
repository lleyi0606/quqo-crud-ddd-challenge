package application

import (
	"products-crud/domain/entity"
	repository "products-crud/domain/repository/product_respository"
	"products-crud/infrastructure/implementations/product"
	base "products-crud/infrastructure/persistences"
)

type productApp struct {
	p *base.Persistence
}

func NewProductApplication(p *base.Persistence) repository.ProductHandlerRepository {
	return &productApp{p}
}

func (u *productApp) AddProduct(user *entity.Product) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)
	return repoProduct.AddProduct(user)
}

func (u *productApp) GetProduct(pdtId uint64) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)
	return repoProduct.GetProduct(pdtId)
}

func (u *productApp) GetProducts() ([]entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)
	return repoProduct.GetProducts()
}

func (u *productApp) UpdateProduct(pdt *entity.ProductUpdate) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)
	return repoProduct.UpdateProduct(pdt)
}

func (u *productApp) DeleteProduct(pdtId uint64) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)
	return repoProduct.DeleteProduct(pdtId)
}

func (u *productApp) SearchProducts(str string) ([]entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)
	return repoProduct.SearchProducts(str)
}
