package opensearch_entity

import "net/http"

// indices
const (
	OpenSearchProductsIndex = "products"
)

type OpenSearch struct {
	Client         *http.Client
	DomainEndpoint string
	Username       string
	Password       string
}
