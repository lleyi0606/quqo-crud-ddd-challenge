package search

import (
	"products-crud/domain/repository/search_repository"
	"products-crud/infrastructure/implementations/search/algolia"
	"products-crud/infrastructure/implementations/search/opensearch"
	base "products-crud/infrastructure/persistences"
)

const (
	Aloglia    = "algolia"
	OpenSearch = "opensearch"
)

func NewSearchRepository(p *base.Persistence, provider string) search_repository.SearchRepository {
	switch provider {
	case Aloglia:
		return algolia.NewAlgoliaRepository(p)
	case OpenSearch:
		return opensearch.NewOpensearchRepository(p)
	default:
		return algolia.NewAlgoliaRepository(p)
	}
}
