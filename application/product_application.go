package application

import (
	"os"
	inventory_entity "products-crud/domain/entity/inventory_entity"
	loggerentity "products-crud/domain/entity/logger_entity"
	entity "products-crud/domain/entity/product_entity"
	repository "products-crud/domain/repository/product_respository"
	"products-crud/infrastructure/implementations/inventory"
	"products-crud/infrastructure/implementations/logger"
	"products-crud/infrastructure/implementations/product"
	"products-crud/infrastructure/implementations/search"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

type productApp struct {
	p *base.Persistence
	c *gin.Context
}

func NewProductApplication(p *base.Persistence, c *gin.Context) repository.ProductHandlerRepository {
	return &productApp{p, c}
}

func (u *productApp) AddProduct(pdt *entity.ProductWithStockAndWarehouse) (*entity.Product, error) {

	info := loggerentity.FunctionInfo{
		FunctionName: "AddProduct",
		Path:         "application/",
		Description:  "Application of AddProduct",
		Body:         pdt,
	}
	logger := logger.NewLoggerRepositories(u.p, u.c, info, "honeycomb", "zap")
	// span := trace.SpanFromContext(*logger.Context)
	// defer span.End()
	defer logger.End()

	repoProduct := product.NewProductRepository(u.p, u.c)

	i := &inventory_entity.Inventory{
		WarehouseID: pdt.WarehouseID,
		Stock:       pdt.Stock,
	}

	repoInventory := inventory.NewInventoryRepository(u.p, nil)
	ivt, err := repoInventory.AddInventory(i)
	if err != nil {
		logger.Error(err.Error(), map[string]interface{}{})
		return nil, err
	}

	p := &entity.Product{
		ProductID:   ivt.ProductID,
		Name:        pdt.Name,
		Description: pdt.Description,
		Price:       pdt.Price,
		CategoryID:  pdt.CategoryID,
		Inventory:   *ivt,
	}

	return repoProduct.AddProduct(p)
}

func (u *productApp) GetProduct(pdtId uint64) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p, nil)
	return repoProduct.GetProduct(pdtId)
}

func (u *productApp) GetProducts() ([]entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p, nil)
	return repoProduct.GetProducts()
}

func (u *productApp) UpdateProduct(pdt *entity.Product) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p, nil)
	return repoProduct.UpdateProduct(pdt)
}

func (u *productApp) DeleteProduct(pdtId uint64) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p, nil)
	return repoProduct.DeleteProduct(pdtId)
}

func (u *productApp) SearchProducts(str string) ([]entity.Product, error) {
	// repoProduct := product.NewProductRepository(u.p, nil)
	// return repoProduct.SearchProducts(str)
	searchTechnology := os.Getenv("SEARCH_TECHNOLOGY")
	repoSearch := search.NewSearchRepository(u.p, searchTechnology)
	return repoSearch.SearchProducts(str)
}
