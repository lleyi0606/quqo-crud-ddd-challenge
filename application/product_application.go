package application

import (
	"encoding/json"
	"log"
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
	logger, endFunc := logger.NewLoggerRepositories(u.p, u.c, info, []string{"Honeycomb", "zap"}, logger.SetNewOtelContext())
	defer endFunc()

	repoProduct := product.NewProductRepository(u.p, u.c)

	i := &inventory_entity.Inventory{
		ProductID:   pdt.ProductID,
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

	// logger.End()
	return repoProduct.AddProduct(p)
}

func (u *productApp) GetProduct(pdtId string) (*entity.Product, error) {
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

func (u *productApp) DeleteProduct(pdtId string) (*entity.Product, error) {
	repoProduct := product.NewProductRepository(u.p, nil)
	return repoProduct.DeleteProduct(pdtId)
}

func (u *productApp) SearchProducts(str string) ([]entity.Product, error) {
	searchTechnology := os.Getenv("SEARCH_TECHNOLOGY")
	repoSearch := search.NewSearchRepository(u.p, u.c, searchTechnology)
	info := loggerentity.FunctionInfo{
		FunctionName: "SearchProducts",
		Path:         "/application/",
		Description:  "Gets ID from search tool and retrieve full data",
		Body:         nil,
	}
	logger, endFunc := logger.NewLoggerRepositories(u.p, u.c, info, []string{"Honeycomb", "zap"})
	defer endFunc()

	resSearch, err := repoSearch.SearchData(str)
	if err != nil {
		logger.Error(err.Error(), map[string]interface{}{})
		return nil, err
	}

	repoProduct := product.NewProductRepository(u.p, u.c)
	var products []entity.Product

	for _, hit := range resSearch {
		// Each hit is a JSON representation of a Product
		jsonBytes, err := json.Marshal(hit)
		if err != nil {
			return nil, err
		}

		log.Print("hit found: ", jsonBytes)

		// Unmarshal the JSON data into a Product struct
		var pdt entity.ProductSearch
		if err := json.Unmarshal(jsonBytes, &pdt); err != nil {
			logger.Error(err.Error()+"location1", map[string]interface{}{"data": pdt})
			return nil, err
		}

		log.Print("pdt search is ", pdt)

		product, err := repoProduct.GetProduct(pdt.ProductID)
		if err != nil {
			logger.Error(err.Error()+"location2", map[string]interface{}{"data": product})
			return nil, err
		}

		// Append the unmarshaled product to the result slice
		products = append(products, *product)
	}

	return products, nil
}
