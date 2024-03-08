package db

import (
	"products-crud/infrastructure/config"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func NewProductSearchDB() (*search.Index, *search.Index, error) {

	app_id := config.Configuration.GetString("algolia.dev.id")
	api_key := config.Configuration.GetString("algolia.dev.key")

	// app_id := os.Getenv("ALGOLIA_APPLICATION_ID")
	// api_key := os.Getenv("ALGOLIA_API_KEY")

	// log.Print("ALGOLIA DB NEWWWWW", aapp_id, aapi_key, "....", app_id, api_key)

	client := search.NewClient(app_id, api_key)
	index := client.InitIndex("products")
	index_inventories := client.InitIndex("inventories")

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
		return nil, nil, err
	}

	_, err = index_inventories.SetSettings(search.Settings{
		AttributesForFaceting: opt.AttributesForFaceting(
			// "author",
			// "filterOnly(isbn)",
			"searchable(name)",
			// "afterDistinct(category)",
			// "afterDistinct(searchable(publisher))",
		),
	})

	if err != nil {
		return nil, nil, err
	}

	return index, index_inventories, nil
}
