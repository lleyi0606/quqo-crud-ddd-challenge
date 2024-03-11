package algolia

import (
	"log"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/repository/search_repository"
	base "products-crud/infrastructure/persistences"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

type algoliaRepo struct {
	p *base.Persistence
	c *gin.Context
}

func (a algoliaRepo) AddProduct(p *entity.Product) error {
	pdtA := entity.SqlProductToProductAloglia(*p)
	_, err := a.p.ProductSearchDb.SaveObject(pdtA)

	if err != nil {
		zap.S().Errorw("Algoria AddProduct ERROR", "error", err)
		return err
	}
	return nil
}

func (a algoliaRepo) DeleteProduct(id string) error {
	_, err := a.p.ProductSearchDb.DeleteObject(id)

	if err != nil {
		zap.S().Errorw("Algoria DeleteProduct ERROR", "error", err)
		return err
	}
	return nil
}

func (a algoliaRepo) SearchData(str string) ([]map[string]interface{}, error) {
	// info := loggerentity.FunctionInfo{
	// 	FunctionName: "SearchProducts",
	// 	Path:         "/infrastructure/handlers/SearchProducts",
	// 	Description:  "Gets query keyword for search",
	// 	Body:         nil,
	// }
	// logger, endFunc := logger.NewLoggerRepositories(a.p, a.c, info, []string{"Honeycomb", "zap"})
	// defer endFunc()

	res, err := a.p.ProductSearchDb.Search(str, opt.AttributesToRetrieve("*"))

	if err != nil {
		// logger.Error(err.Error(), map[string]interface{}{"data": str})
		return nil, err
	}

	return res.Hits, nil

	// var products []entity.Product

	// for _, hit := range res.Hits {
	// 	// Each hit is a JSON representation of a Product
	// 	jsonBytes, err := json.Marshal(hit)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	// Unmarshal the JSON data into a Product struct
	// 	var product entity.Product
	// 	if err := json.Unmarshal(jsonBytes, &product); err != nil {
	// 		return nil, err
	// 	}

	// 	// Append the unmarshaled product to the result slice
	// 	products = append(products, product)
	// }

	// return products, nil
}

func (a algoliaRepo) UpdateProduct(p *entity.Product) error {

	product := entity.ProductSearch{
		Name:      p.Name,
		ProductID: p.ProductID,
		// ObjectID:  p.ProductID,
	}

	_, err := a.p.ProductSearchDb.PartialUpdateObject(product)
	log.Print(p)
	if err != nil {
		zap.S().Errorw("Algolia UpdateProduct error", "error", err, "product", p)
		return err
	}
	return nil
}

// func (a algoliaRepo) AddInventory(i *inventory_entity.Inventory) error {
// 	ivtA := inventory_entity.SqlInventorytoInventoryAlgolia(*i)
// 	_, err := a.p.InventoryAlgoliaDb.SaveObject(ivtA)

// 	if err != nil {
// 		zap.S().Errorw("Algoria AddInventory ERROR", "error", err)
// 		return err
// 	}
// 	return nil
// }

// func (a algoliaRepo) SearchInventories(str string) ([]inventory_entity.Inventory, error) {
// 	res, err := a.p.InventoryAlgoliaDb.Search(str, opt.AttributesToRetrieve("*"))

// 	if err != nil {
// 		return nil, err
// 	}

// 	var inventories []inventory_entity.Inventory

// 	for _, hit := range res.Hits {
// 		// Each hit is a JSON representation of a Product
// 		jsonBytes, err := json.Marshal(hit)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Unmarshal the JSON data into a Product struct
// 		var inventory inventory_entity.Inventory
// 		if err := json.Unmarshal(jsonBytes, &inventory); err != nil {
// 			return nil, err
// 		}

// 		// Append the unmarshaled product to the result slice
// 		inventories = append(inventories, inventory)
// 	}

// 	return inventories, nil
// }

// func (a algoliaRepo) DeleteInventory(id uint64) error {
// 	_, err := a.p.InventoryAlgoliaDb.DeleteObject(strconv.FormatUint(id, 10))

// 	if err != nil {
// 		zap.S().Errorw("Algoria DeleteInventory ERROR", "error", err)
// 		return err
// 	}
// 	return nil
// }

// func (a algoliaRepo) UpdateInventory(i *inventory_entity.Inventory) error {

// 	inventory := inventory_entity.InventoryAlgolia{
// 		Inventory: *i,
// 		ObjectID:  i.InventoryID,
// 	}

// 	_, err := a.p.ProductSearchDb.PartialUpdateObject(inventory)
// 	log.Print(i)
// 	if err != nil {
// 		zap.S().Errorw("Algolia UpdateInventory error", "error", err, "inventory", i)
// 		return err
// 	}
// 	return nil
// }

func NewAlgoliaRepository(p *base.Persistence, c *gin.Context) search_repository.SearchRepository {
	return &algoliaRepo{p, c}
}
