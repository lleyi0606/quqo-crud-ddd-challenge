package search

import (
	"products-crud/domain/repository/search_repository"
	"products-crud/infrastructure/implementations/search/algolia"
	base "products-crud/infrastructure/persistences"
)

const (
	Aloglia = "algolia"
)

func NewSearchRepository(p *base.Persistence, provider string) search_repository.SearchRepository {
	switch provider {
	case Aloglia:
		return algolia.NewAlgoliaRepository(p)
	default:
		return algolia.NewAlgoliaRepository(p)
	}
}
