package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(r *gin.RouterGroup, p *base.Persistence) {

	inventoryHandler := handlers.NewInventoryController(p)

	r.POST("/inventories", inventoryHandler.AddInventory)
	r.POST("/inventories/all", inventoryHandler.AddInventories)
	r.GET("/inventories", inventoryHandler.GetInventories)
	r.GET("/inventories/:id", inventoryHandler.GetInventory)
	r.PUT("/inventories/:id", inventoryHandler.UpdateInventory)
	r.DELETE("/inventories/:id", inventoryHandler.DeleteInventory)
	r.GET("/inventories/search", inventoryHandler.SearchInventory)

}
