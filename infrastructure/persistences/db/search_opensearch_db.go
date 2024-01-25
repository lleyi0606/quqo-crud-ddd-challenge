package db

import (
	"net/http"

	"products-crud/domain/entity/opensearch_entity"
	"products-crud/infrastructure/config"
)

func NewProductOpenSearchDB() (*opensearch_entity.OpenSearch, error) {

	openS := &opensearch_entity.OpenSearch{
		Client: &http.Client{},
		// DomainEndpoint: os.Getenv("AWS_DOMAIN_ENDPOINT"),
		// Username:       os.Getenv("AWS_USER"),
		// Password:       os.Getenv("AWS_PASSWORD"),
		DomainEndpoint: config.Configuration.GetString("awsOpensearch.domainEndpoint"),
		Username:       config.Configuration.GetString("awsOpensearch.opensearch.dev.user"),
		Password:       config.Configuration.GetString("awsOpensearch.opensearch.dev.pass"),
	}

	return openS, nil
}
