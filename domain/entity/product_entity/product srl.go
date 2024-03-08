package entity

func SqlProductToProductAloglia(p Product) ProductSearch {
	var product ProductSearch

	product.Name = p.Name
	product.ProductID = p.ProductID
	// product.ObjectID = p.ProductID
	return product
}
