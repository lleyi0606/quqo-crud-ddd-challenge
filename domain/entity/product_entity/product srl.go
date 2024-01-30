package entity

func SqlProductToProductAloglia(p Product) ProductAlgolia {
	var product ProductAlgolia

	product.ProductID = p.ProductID
	product.Name = p.Name
	product.Description = p.Description
	product.Price = p.Price
	product.Category = p.Category
	product.ObjectID = p.ProductID
	return product
}
