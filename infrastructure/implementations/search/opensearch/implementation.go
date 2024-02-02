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

	// openS := o.p.SearchOpenSearchDb

	pdt, err := json.Marshal(p)
	// // documentID := p.ProductID
	documentID := fmt.Sprint(p.ProductID)

	// log.Print(documentID, pdt)

	// mapping := `{
	// 	"mappings" : {
	// 		"properties" :  {
	// 			"counter" : {
	// 				"type" : "unsigned_long"
	// 			}
	// 		}
	// 	}
	// }`
	// requestBody := mapping + string(pdt)

	// // url := fmt.Sprintf("%s/%s/_doc/%v", openS.DomainEndpoint, opensearch_entity.OpenSearchProductsIndex, documentID)

	// req, err := http.NewRequest("PUT", url, strings.NewReader(requestBody))
	// if err != nil {
	// 	fmt.Println("Error creating request:", err)
	// 	return err
	// }

	// req.SetBasicAuth(openS.Username, openS.Password)
	// req.Header.Set("Content-Type", "application/json")

	// resp, err := openS.Client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending request:", err)
	// 	return err
	// }
	// defer resp.Body.Close()

	// Create an IndexRequest with the mapping included in the request body GOOOOOOD
	// mapping := `
	// {
	//   "mappings": {
	// 	"properties": {
	// 	  "productID": { "type": "unsigned_long" },
	// 	  "name": { "type": "text" },
	// 	  "description": { "type": "text" },
	// 	  "price": { "type": "double" },
	// 	  "category": { "type": "keyword" },
	// 	  "warehouseID": { "type": "unsigned_long" },
	// 	  "stock": { "type": "integer" }
	// 	}
	//   }
	// }`

	// req_indice := opensearchapi.IndicesCreateRequest{
	// 	Index: opensearch_entity.OpenSearchProductsIndex,
	// 	Body:  strings.NewReader(mapping),
	// }

	// createIndexResponse, err := req_indice.Do(context.Background(), openS)
	// if err != nil {
	// 	fmt.Println("Failed to create index with mapping:", err)
	// 	return err
	// }
	// defer createIndexResponse.Body.Close()

	// // Check the response
	// if createIndexResponse.IsError() {
	// 	fmt.Printf("Error creating index: %s\n", createIndexResponse.Status())
	// 	return fmt.Errorf("failed to create index: %s", createIndexResponse.Status())
	// }

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

	// content := fmt.Sprintf(`{
	// 	"size": 100,
	// 	"_source": {
	// 	  "includes": ["name", "description"]
	// 	},
	// 	"query": {
	// 	  "neural": {
	// 		"name_v": {
	// 		  "query_text": "%s",
	// 		  "model_id": "dZGTYo0BqHh94dza0aSQ",
	// 		  "k": 100
	// 		}
	// 	  }
	// 	}
	//   }`, str)

	contentReader := strings.NewReader(content)

	search := opensearchapi.SearchRequest{
		Index: []string{opensearch_entity.OpenSearchProductsIndex},
		// Index: []string{"products_ml"},
		Body: contentReader,
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

// func (o opensearchRepo) UpdateProduct(p *entity.Product) error {
// 	openS := o.p.SearchOpenSearchDb

// 	pdt, err := json.Marshal(p)
// 	documentID := p.ProductID

// 	url := fmt.Sprintf("%s/%s/_update/%d", openS.DomainEndpoint, opensearch_entity.OpenSearchProductsIndex, documentID)

// 	req, err := http.NewRequest("POST", url, strings.NewReader(string(pdt)))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return err
// 	}

// 	req.SetBasicAuth(openS.Username, openS.Password)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := openS.Client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Check the response
// 	if resp.StatusCode == 201 {
// 		fmt.Println("Document updated successfully.", documentID)
// 	} else {
// 		// Read the response body
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println("Error reading response body:", err)
// 			return err
// 		}

// 		// Convert the body to a string and print it
// 		fmt.Printf("Failed to update document. Status code: %d\n Body: %s\n", resp.StatusCode, string(body))
// 	}

// 	return nil

// }

func NewOpensearchRepository(p *base.Persistence) search_repository.SearchRepository {
	return &opensearchRepo{p}
}
