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

func SqlProductToProductAloglia(p Product) ProductAlgolia {
	var product ProductAlgolia

	product.ID = p.ID
	product.Name = p.Name
	if p.Description != nil {
		product.Description = *p.Description
	}
	product.Price = p.Price
	product.Category = p.Category
	product.Stock = p.Stock
	product.Image = p.Image
	// product.ObjectID = strconv.FormatUint(p.ID, 10)
	product.ObjectID = p.ID
	return product
}
