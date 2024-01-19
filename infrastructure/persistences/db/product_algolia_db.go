package db

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func NewProductAlgoliaDB() (*search.Index, error) {

	app_id := os.Getenv("ALGOLIA_APPLICATION_ID")
	api_key := os.Getenv("ALGOLIA_API_KEY")

	client := search.NewClient(app_id, api_key)
	index := client.InitIndex("products")

	_, err := index.SetSettings(search.Settings{
		AttributesForFaceting: opt.AttributesForFaceting(
			// "author",
			// "filterOnly(isbn)",
			"searchable(name)",
			// "afterDistinct(category)",
			// "afterDistinct(searchable(publisher))",
		),
	})

	if err != nil {
		return nil, err
	}

	return index, nil
}
