package redis_entity

import "time"

type BlacklistedToken struct {
	Token       string `json:"token"`
	Reason      string `json:"reason,omitempty"`
	RevokedByID string `json:"revoked_by_id,omitempty"`
}

// keys
const (
	RedisProducts      = "PRODUCTS_ENABLED_LIST_DATA"
	RedisProductData   = "PRODUCT_"
	RedisInventories   = "INVENTORIES_ENABLED_LIST_DATA"
	RedisInventoryData = "INVENTORY_"
	RedisImages        = "IMAGES_ENABLED_LIST_DATA"
	RedisImageData     = "IMAGE_"
	RedisJWTData       = "JWT_"
)

// expirations
const (
	RedisExpirationGlobal = time.Minute * 2
	RedisExpirationJwt    = time.Hour * 240
)
