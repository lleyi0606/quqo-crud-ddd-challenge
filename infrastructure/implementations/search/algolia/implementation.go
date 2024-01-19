package algolia

import (
	"encoding/json"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/repository/search_repository"
	base "products-crud/infrastructure/persistences"

	"go.uber.org/zap"
)

type algoliaRepo struct {
	p *base.Persistence
}

func (a algoliaRepo) AddProduct(p *entity.Product) error {
	_, err := a.p.ProductAlgoliaDb.SaveObject(p)

	if err != nil {
		zap.S().Errorw("2. Redis SetKey ERROR", "error", err)
		return err
	}
	return nil
}

func (a algoliaRepo) SearchProducts(str string) ([]entity.Product, error) {
	res, err := a.p.ProductAlgoliaDb.SearchForFacetValues("name", str)
	if err != nil {
		return nil, err
	}
	var products []entity.Product

	for _, hit := range res.FacetHits {
		var product entity.Product

		// Assuming each hit is a JSON representation of a Product
		jsonBytes, err := json.Marshal(hit)
		if err != nil {
			return nil, err
		}

		// Unmarshal the JSON data into a Product struct
		if err := json.Unmarshal(jsonBytes, &product); err != nil {
			return nil, err
		}

		// Append the unmarshaled product to the result slice
		products = append(products, product)
	}

	return products, nil
}

func NewAlgoliaRepository(p *base.Persistence) search_repository.SearchRepository {
	return &algoliaRepo{p}
}
