package db

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func NewProductAlgoliaDB() (*search.Index, error) {

	app_id := os.Getenv("ALGOLIA_APPLICATION_ID")
	api_key := os.Getenv("ALGOLIA_APPLICATION_ID")

	client := search.NewClient(app_id, api_key)
	index := client.InitIndex("products")

	return index, nil
}
