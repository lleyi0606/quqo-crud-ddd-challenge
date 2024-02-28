package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func OrderedItemRoutes(r *gin.RouterGroup, p *base.Persistence) {

	OrderedItemHandler := handlers.NewOrderedItemController(p)

	r.GET("/ordereditems", OrderedItemHandler.GetOrderedItems)
	r.GET("/ordereditems/:id", OrderedItemHandler.GetOrderedItemsByOrderId)

}
