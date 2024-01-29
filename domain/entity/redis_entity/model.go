package redis_entity

import "time"

// keys
const (
	RedisProducts      = "PRODUCTS_ENABLED_LIST_DATA"
	RedisProductData   = "PRODUCT_"
	RedisInventories   = "INVENTORIES_ENABLED_LIST_DATA"
	RedisInventoryData = "INVENTORY_"
	RedisImages        = "IMAGES_ENABLED_LIST_DATA"
	RedisImageData     = "IMAGE_"
)

// expirations
const (
	RedisExpirationGlobal = time.Minute * 2
)
