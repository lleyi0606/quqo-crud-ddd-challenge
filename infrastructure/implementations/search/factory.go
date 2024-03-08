package search

import (
	"products-crud/domain/repository/search_repository"
	"products-crud/infrastructure/implementations/search/algolia"

	"products-crud/infrastructure/implementations/search/opensearch"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

const (
	Algolia    = "algolia"
	OpenSearch = "opensearch"
)

func NewSearchRepository(p *base.Persistence, c *gin.Context, provider string) search_repository.SearchRepository {
	switch provider {
	case Algolia:
		return algolia.NewAlgoliaRepository(p, c)
	case OpenSearch:
		return opensearch.NewOpensearchRepository(p)
	default:
		return algolia.NewAlgoliaRepository(p, c)
	}
}
