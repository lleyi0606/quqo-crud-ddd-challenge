package db

import (
	"net/http"
	"os"
	"products-crud/domain/entity/opensearch_entity"
)

func NewProductOpenSearchDB() (*opensearch_entity.OpenSearch, error) {

	openS := &opensearch_entity.OpenSearch{
		Client:         &http.Client{},
		DomainEndpoint: os.Getenv("AWS_DOMAIN_ENDPOINT"),
		Username:       os.Getenv("AWS_USER"),
		Password:       os.Getenv("AWS_PASSWORD"),
	}

	return openS, nil
}
