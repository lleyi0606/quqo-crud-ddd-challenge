package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup, p *base.Persistence) {

	productHandler := handlers.NewProductController(p)

	r.POST("/products", productHandler.AddProduct)
	r.POST("/products/all", productHandler.AddProducts)
	r.GET("/products", productHandler.GetProducts)
	r.GET("/user/products", productHandler.GetProductsUser)
	r.GET("/products/:id", productHandler.GetProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)
	r.GET("/products/search", productHandler.SearchProducts)

}
