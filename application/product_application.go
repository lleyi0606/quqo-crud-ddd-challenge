package application

import (
	inventory_entity "products-crud/domain/entity/inventory_entity"
	entity "products-crud/domain/entity/product_entity"
	repository "products-crud/domain/repository/product_respository"
	"products-crud/infrastructure/implementations/inventory"
	"products-crud/infrastructure/implementations/product"
	base "products-crud/infrastructure/persistences"
)

type productApp struct {
	p *base.Persistence
}

func NewProductApplication(p *base.Persistence) repository.ProductHandlerRepository {
	return &productApp{p}
}

func (u *productApp) AddProduct(pdt *entity.ProductWithStockAndWarehouse) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)

	p := &entity.Product{
		ProductID:   pdt.ProductID,
		Name:        pdt.Name,
		Description: pdt.Description,
		Price:       pdt.Price,
		Category:    pdt.Category,
	}

	i := &inventory_entity.Inventory{
		ProductID:   pdt.ProductID,
		WarehouseID: pdt.WarehouseID,
		Stock:       pdt.Stock,
	}

	repoInventory := inventory.NewInventoryRepository(u.p)
	_, err := repoInventory.AddInventory(i)
	if err != nil {
		return nil, err
	}

	return repoProduct.AddProduct(p)
}

func (u *productApp) GetProduct(pdtId uint64) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)
	return repoProduct.GetProduct(pdtId)
}

func (u *productApp) GetProducts() ([]entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p)
	return repoProduct.GetProducts()
}

func (u *productApp) UpdateProduct(pdt *entity.Product) (*entity.Product, error) {
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
