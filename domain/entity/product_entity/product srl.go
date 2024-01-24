package entity

func SqlProductToProductAloglia(p Product) ProductAlgolia {
	var product ProductAlgolia

	product.ID = p.ID
	product.Name = p.Name
	if p.Description != nil {
		product.Description = p.Description
	}
	product.Price = p.Price
	product.Category = p.Category
	product.Stock = p.Stock
	product.Image = p.Image
	product.ObjectID = p.ID
	return product
}
