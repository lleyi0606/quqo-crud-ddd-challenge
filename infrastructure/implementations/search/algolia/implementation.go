package algolia

import (
	"encoding/json"
	"log"
	inventory_entity "products-crud/domain/entity/inventory_entity"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/repository/search_repository"
	base "products-crud/infrastructure/persistences"
	"strconv"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"

	"go.uber.org/zap"
)

type algoliaRepo struct {
	p *base.Persistence
}

func (a algoliaRepo) AddProduct(p *entity.Product) error {
	pdtA := entity.SqlProductToProductAloglia(*p)
	_, err := a.p.ProductAlgoliaDb.SaveObject(pdtA)

	if err != nil {
		zap.S().Errorw("Algoria AddProduct ERROR", "error", err)
		return err
	}
	return nil
}

func (a algoliaRepo) DeleteProduct(id uint64) error {
	_, err := a.p.ProductAlgoliaDb.DeleteObject(strconv.FormatUint(id, 10))

	if err != nil {
		zap.S().Errorw("Algoria DeleteProduct ERROR", "error", err)
		return err
	}
	return nil
}

func (a algoliaRepo) SearchProducts(str string) ([]entity.Product, error) {
	res, err := a.p.ProductAlgoliaDb.Search(str, opt.AttributesToRetrieve("*"))

	if err != nil {
		return nil, err
	}

	var products []entity.Product

	for _, hit := range res.Hits {
		// Each hit is a JSON representation of a Product
		jsonBytes, err := json.Marshal(hit)
		if err != nil {
			return nil, err
		}

		// Unmarshal the JSON data into a Product struct
		var product entity.Product
		if err := json.Unmarshal(jsonBytes, &product); err != nil {
			return nil, err
		}

		// Append the unmarshaled product to the result slice
		products = append(products, product)
	}

	return products, nil
}

func (a algoliaRepo) UpdateProduct(p *entity.Product) error {

	product := entity.ProductAlgolia{
		Product:  *p,
		ObjectID: p.ID, // Convert ID to string
	}

	_, err := a.p.ProductAlgoliaDb.PartialUpdateObject(product)
	log.Print(p)
	if err != nil {
		zap.S().Errorw("Algolia UpdateProduct error", "error", err, "product", p)
		return err
	}
	return nil
}

func (a algoliaRepo) AddInventory(i *inventory_entity.Inventory) error {
	ivtA := inventory_entity.SqlInventorytoInventoryAlgolia(*i)
	_, err := a.p.InventoryAlgoliaDb.SaveObject(ivtA)

	if err != nil {
		zap.S().Errorw("Algoria AddInventory ERROR", "error", err)
		return err
	}
	return nil
}

func (a algoliaRepo) SearchInventories(str string) ([]inventory_entity.Inventory, error) {
	res, err := a.p.InventoryAlgoliaDb.Search(str, opt.AttributesToRetrieve("*"))

	if err != nil {
		return nil, err
	}

	var inventories []inventory_entity.Inventory

	for _, hit := range res.Hits {
		// Each hit is a JSON representation of a Product
		jsonBytes, err := json.Marshal(hit)
		if err != nil {
			return nil, err
		}

		// Unmarshal the JSON data into a Product struct
		var inventory inventory_entity.Inventory
		if err := json.Unmarshal(jsonBytes, &inventory); err != nil {
			return nil, err
		}

		// Append the unmarshaled product to the result slice
		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

func (a algoliaRepo) DeleteInventory(id uint64) error {
	_, err := a.p.InventoryAlgoliaDb.DeleteObject(strconv.FormatUint(id, 10))

	if err != nil {
		zap.S().Errorw("Algoria DeleteInventory ERROR", "error", err)
		return err
	}
	return nil
}

func (a algoliaRepo) UpdateInventory(i *inventory_entity.Inventory) error {

	inventory := inventory_entity.InventoryAlgolia{
		Inventory: *i,
		ObjectID:  i.InventoryID,
	}

	_, err := a.p.ProductAlgoliaDb.PartialUpdateObject(inventory)
	log.Print(i)
	if err != nil {
		zap.S().Errorw("Algolia UpdateInventory error", "error", err, "inventory", i)
		return err
	}
	return nil
}

func NewAlgoliaRepository(p *base.Persistence) search_repository.SearchRepository {
	return &algoliaRepo{p}
}
