package algolia

import (
	"encoding/json"
	"log"
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

	// var product entity.ProductAlgolia

	// product.ID = p.ID
	// product.Name = p.Name
	// product.Description = p.Description
	// product.Price = p.Price
	// product.Category = p.Category
	// product.Stock = p.Stock
	// product.Image = p.Image
	// product.ObjectID = p.ID

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

func NewAlgoliaRepository(p *base.Persistence) search_repository.SearchRepository {
	return &algoliaRepo{p}
}
