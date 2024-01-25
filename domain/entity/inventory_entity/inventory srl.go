package entity

func SqlInventorytoInventoryAlgolia(i Inventory) InventoryAlgolia {
	var ivt InventoryAlgolia

	ivt.ProductID = i.ProductID
	ivt.WarehouseID = i.WarehouseID
	ivt.Stock = i.Stock
	ivt.ObjectID = i.ProductID
	return ivt
}
