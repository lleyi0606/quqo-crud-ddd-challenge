package db

import (
	"log"
	"products-crud/infrastructure/config"

	storage_go "github.com/supabase-community/storage-go"
)

// ImageSupabaseDB represents the Supabase storage client.
type ImageSupabaseDB struct {
	Client *storage_go.Client
}

// NewImageSupabaseDB creates a new instance of the Supabase storage client.
func NewImageSupabaseDB() (*ImageSupabaseDB, error) {
	storageURL := config.Configuration.GetString("supabase.dev.url") + "/storage/v1"
	// storageURL := config.Configuration.GetString("supabase.dev.url")
	storageKey := config.Configuration.GetString("supabase.dev.key")

	log.Print(storageKey + "    " + storageURL)

	storageClient := storage_go.NewClient(storageURL, storageKey, nil)

	// You may want to check for errors during client creation
	if storageClient == nil {
		return nil, &storage_go.StorageError{} // Replace someError with appropriate error handling
	}

	// _, err := storageClient.CreateBucket("images", storage_go.BucketOptions{
	// 	Public: true,
	// })

	// if err != nil {
	// 	return nil, err
	// }

	log.Print("end of image_supabase_db file")

	return &ImageSupabaseDB{
		Client: storageClient,
	}, nil
}
