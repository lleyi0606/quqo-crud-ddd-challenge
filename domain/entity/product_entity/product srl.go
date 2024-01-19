package entity

import (
	"encoding/json"
	"fmt"
)

func SqlProductRToProductForUpdate(p ProductToReceive, id uint64) ProductUpdate {
	// Convert ProductToReceive to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		// returns a default Product
		return ProductUpdate{}
	}

	// Convert JSON to Product
	var product ProductUpdate
	err = json.Unmarshal(jsonData, &product)
	if err != nil {
		// returns a default Product
		return ProductUpdate{}
	}

	product.ID = id
	fmt.Println(product)

	return product
}