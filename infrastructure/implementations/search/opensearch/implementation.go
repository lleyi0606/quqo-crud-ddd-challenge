package opensearch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"products-crud/domain/entity/opensearch_entity"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/repository/search_repository"
	base "products-crud/infrastructure/persistences"
	"strconv"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"go.uber.org/zap"
)

type opensearchRepo struct {
	p *base.Persistence
}

func (o opensearchRepo) AddProduct(p *entity.Product) error {

	openS := o.p.SearchOpenSearchDb

	pdt, err := json.Marshal(p)
	documentID := p.ID

	url := fmt.Sprintf("%s/%s/_doc/%d", openS.DomainEndpoint, opensearch_entity.OpenSearchProductsIndex, documentID)

	req, err := http.NewRequest("PUT", url, strings.NewReader(string(pdt)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.SetBasicAuth(openS.Username, openS.Password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := openS.Client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode == 201 {
		fmt.Println("Document indexed successfully.", documentID)
	} else {
		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return err
		}

		// Convert the body to a string and print it
		fmt.Printf("Failed to index document. Status code: %d\n Body: %s\n", resp.StatusCode, string(body))
	}

	return nil
}

func (o opensearchRepo) DeleteProduct(id uint64) error {
	_, err := o.p.ProductAlgoliaDb.DeleteObject(strconv.FormatUint(id, 10))

	if err != nil {
		zap.S().Errorw("Algoria DeleteProduct ERROR", "error", err)
		return err
	}
	return nil
}

func (o opensearchRepo) SearchProducts(str string) ([]entity.Product, error) {
	res, err := o.p.ProductAlgoliaDb.Search(str, opt.AttributesToRetrieve("*"))

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

func (o opensearchRepo) UpdateProduct(p *entity.Product) error {

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

	_, err := o.p.ProductAlgoliaDb.PartialUpdateObject(product)
	log.Print(p)
	if err != nil {
		zap.S().Errorw("Algolia UpdateProduct error", "error", err, "product", p)
		return err
	}
	return nil
}

func NewOpensearchRepository(p *base.Persistence) search_repository.SearchRepository {
	return &opensearchRepo{p}
}
