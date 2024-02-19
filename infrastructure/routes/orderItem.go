package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func OrderedItemRoutes(r *gin.RouterGroup, p *base.Persistence) {

	OrderedItemHandler := handlers.NewOrderedItemController(p)

	r.GET("/orderedItems", OrderedItemHandler.GetOrderedItems)
	r.GET("/orderedItems/:id", OrderedItemHandler.GetOrderedItem)
	r.PUT("/orderedItems/:id", OrderedItemHandler.UpdateOrderedItem)
	r.DELETE("/orderedItems/:id", OrderedItemHandler.DeleteOrderedItem)

}
