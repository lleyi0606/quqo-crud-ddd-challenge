package opensearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"products-crud/domain/entity/opensearch_entity"
	entity "products-crud/domain/entity/product_entity"
	"products-crud/domain/repository/search_repository"
	base "products-crud/infrastructure/persistences"
	"strings"
)

type opensearchRepo struct {
	p *base.Persistence
}

func (o opensearchRepo) AddProduct(p *entity.Product) error {

	openS := o.p.SearchOpenSearchDb

	pdt, err := json.Marshal(p)
	// documentID := p.ProductID
	documentID := fmt.Sprint(p.ProductID)

	log.Print(documentID, pdt)

	mapping := `{
		"mappings" : {
		  "properties" :  {
			"counter" : {
			  "type" : "unsigned_long"
			}
		  }
		}
	}`

	requestBody := mapping + string(pdt)

	url := fmt.Sprintf("%s/%s/_doc/%v", openS.DomainEndpoint, opensearch_entity.OpenSearchProductsIndex, documentID)

	req, err := http.NewRequest("PUT", url, strings.NewReader(requestBody))
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

	openS := o.p.SearchOpenSearchDb

	url := fmt.Sprintf("%s/%s/_delete/%d", openS.DomainEndpoint, opensearch_entity.OpenSearchProductsIndex, id)

	req, err := http.NewRequest("DEL", url, nil)
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
		fmt.Println("Document deleted successfully.", id)
	} else {
		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return err
		}

		// Convert the body to a string and print it
		fmt.Printf("Failed to delete document. Status code: %d\n Body: %s\n", resp.StatusCode, string(body))
	}

	return nil
}

func (o opensearchRepo) SearchProducts(str string) ([]entity.Product, error) {
	openS := o.p.SearchOpenSearchDb

	// url := fmt.Sprintf("%s/%s/_search?q=%s&pretty=true", openS.DomainEndpoint, opensearch_entity.OpenSearchProductsIndex, str)

	url := fmt.Sprintf("%s/%s/_search", openS.DomainEndpoint, opensearch_entity.OpenSearchProductsIndex)

	// Build the request body
	query := fmt.Sprintf(`{
		"query": {
			"fuzzy": {
			"name": {
				"value": "%s",
				"fuzziness": "2",
				"max_expansions": 40,
				"prefix_length": 0,
				"transpositions": true,
				"rewrite": "constant_score"
			}
			}
		}
		}`, str)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(openS.Username, openS.Password)

	resp, err := openS.Client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
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

func (o opensearchRepo) UpdateProduct(p *entity.Product) error {
	openS := o.p.SearchOpenSearchDb

	pdt, err := json.Marshal(p)
	documentID := p.ProductID

	url := fmt.Sprintf("%s/%s/_update/%d", openS.DomainEndpoint, opensearch_entity.OpenSearchProductsIndex, documentID)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(pdt)))
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
		fmt.Println("Document updated successfully.", documentID)
	} else {
		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return err
		}

		// Convert the body to a string and print it
		fmt.Printf("Failed to update document. Status code: %d\n Body: %s\n", resp.StatusCode, string(body))
	}

	return nil

}

func NewOpensearchRepository(p *base.Persistence) search_repository.SearchRepository {
	return &opensearchRepo{p}
}
