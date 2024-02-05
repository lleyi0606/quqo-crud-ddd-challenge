package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.RouterGroup, p *base.Persistence) {

	CustomerHandler := handlers.NewCustomerController(p)

	r.POST("/customers", CustomerHandler.AddCustomer)
	r.GET("/customers/:id", CustomerHandler.GetCustomer)
	r.PUT("/customers/:id", CustomerHandler.UpdateCustomer)
	r.DELETE("/customers/:id", CustomerHandler.DeleteCustomer)

}
