package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"products-crud/domain/entity/opensearch_entity"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/repository/search_repository"
	base "products-crud/infrastructure/persistences"
	"strings"

	opensearchapi "github.com/opensearch-project/opensearch-go/opensearchapi"
)

type opensearchRepo struct {
	p *base.Persistence
}

func (o opensearchRepo) AddProduct(p *entity.Product) error {

	pdt, err := json.Marshal(p)
	documentID := fmt.Sprint(p.ProductID)

	req := opensearchapi.IndexRequest{
		Index:      opensearch_entity.OpenSearchProductsIndex,
		DocumentID: documentID,
		Body:       strings.NewReader(string(pdt)),
	}
	insertResponse, err := req.Do(context.Background(), o.p.SearchOpenSearchDb)
	if err != nil {
		fmt.Println("failed to insert document ", err)
		os.Exit(1)
	}

	fmt.Println("Inserting a document", string(pdt))
	fmt.Println(insertResponse)
	defer insertResponse.Body.Close()

	// Check the response
	if insertResponse.StatusCode == 201 {
		fmt.Println("Document indexed successfully.", documentID)
	} else {
		// Read the response body
		body, err := io.ReadAll(insertResponse.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return err
		}

		// Convert the body to a string and print it
		fmt.Printf("Failed to index document. Status code: %d\n Body: %s\n", insertResponse.StatusCode, string(body))
	}

	return nil
}

func (o opensearchRepo) DeleteProduct(id uint64) error {

	delete := opensearchapi.DeleteRequest{
		Index:      opensearch_entity.OpenSearchProductsIndex,
		DocumentID: fmt.Sprint(id),
	}

	deleteResponse, err := delete.Do(context.Background(), o.p.SearchOpenSearchDb)
	if err != nil {
		fmt.Println("failed to delete document ", err)
		os.Exit(1)
	}
	fmt.Println("Deleting a document")
	fmt.Println(deleteResponse)
	defer deleteResponse.Body.Close()
	return nil
}

func (o opensearchRepo) SearchProducts(str string) ([]entity.Product, error) {
	log.Print("keyword: ", str)
	// Search for the document.
	content := fmt.Sprintf(`{
		"size": 10,
		"query": {
			"bool": {
				"should": [
					{
						"match_phrase": {
							"name": {
								"query": "%s",
								"boost": 2,
								"slop": 2
							}
						}
					},
					{
						"multi_match": {
							"query": "%s",
							"fields": ["name^3", "description"], 
							"fuzziness": "AUTO"
						}
					}
				],
				"minimum_should_match": 1
			}
		},
		"sort": [
			{
				"_score": {
					"order": "desc"
				}
			}
		]
	}
	
	
	  `, str, str)

	contentReader := strings.NewReader(content)

	search := opensearchapi.SearchRequest{
		Index: []string{opensearch_entity.OpenSearchProductsIndex},
		Body:  contentReader,
	}

	resp, err := search.Do(context.Background(), o.p.SearchOpenSearchDb)
	if err != nil {
		fmt.Println("failed to search document ", err)
		os.Exit(1)
	}
	fmt.Println("Searching for a document")
	fmt.Println(resp)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	// Check the response
	if resp.StatusCode == 200 {
		fmt.Println("Document searched successfully.")
	} else {
		// Read the response body
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil, err
		}

		// Convert the body to a string and print it
		fmt.Printf("Failed to search documents. Status code: %d\n Body: %s\n", resp.StatusCode, string(body))
	}

	var products []entity.Product

	// Parse the response JSON
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	// Extract hits from the response
	hits, ok := response["hits"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("hits not found in response")
	}

	hitsList, ok := hits["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("hits list not found in response")
	}

	// Check if hits list is empty and return an empty array
	if len(hitsList) == 0 {
		return []entity.Product{}, nil
	}

	// Iterate through hits and extract products
	for _, hit := range hitsList {
		hitMap, ok := hit.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("hit is not a map")
		}

		source, ok := hitMap["_source"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("_source not found in hit")
		}

		// Convert _source to JSON
		jsonBytes, err := json.Marshal(source)
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

func NewOpensearchRepository(p *base.Persistence) search_repository.SearchRepository {
	return &opensearchRepo{p}
}
