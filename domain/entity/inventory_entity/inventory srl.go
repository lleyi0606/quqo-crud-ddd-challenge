package entity

func SqlInventorytoInventoryAlgolia(i Inventory) InventoryAlgolia {
	var ivt InventoryAlgolia

	ivt.InventoryID = i.InventoryID
	ivt.Name = i.Name
	ivt.Address = i.Address
	ivt.ObjectID = i.InventoryID
	return ivt
}
