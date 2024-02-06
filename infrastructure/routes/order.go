package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.RouterGroup, p *base.Persistence) {

	OrderHandler := handlers.NewOrderController(p)

	r.POST("/orders", OrderHandler.AddOrder)
	r.GET("/orders/:id", OrderHandler.GetOrder)
	r.PUT("/orders/:id", OrderHandler.UpdateOrder)
	r.DELETE("/orders/:id", OrderHandler.DeleteOrder)

}
