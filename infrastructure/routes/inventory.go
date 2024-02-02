package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(r *gin.RouterGroup, p *base.Persistence) {

	inventoryHandler := handlers.NewInventoryController(p)

	r.GET("/inventories/:id", inventoryHandler.GetInventory)
	r.PUT("/inventories/:id", inventoryHandler.UpdateStock)

}
