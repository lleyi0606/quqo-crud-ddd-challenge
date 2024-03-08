package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func WarehouseRoutes(r *gin.RouterGroup, p *base.Persistence) {

	WarehouseHandler := handlers.NewWarehouseController(p)

	r.POST("/warehouses", WarehouseHandler.AddWarehouse)
	r.GET("/warehouses/:id", WarehouseHandler.GetWarehouse)
	r.GET("/warehouses", WarehouseHandler.GetWarehouses)
	r.PUT("/warehouses/:id", WarehouseHandler.UpdateWarehouse)
	r.DELETE("/warehouses/:id", WarehouseHandler.DeleteWarehouse)

}
