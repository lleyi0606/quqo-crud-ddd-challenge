package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup, p *base.Persistence) {

	CategoryHandler := handlers.NewCategoryController(p)

	r.POST("/categories", CategoryHandler.AddCategory)
	r.GET("/categories/:id", CategoryHandler.GetCategory)
	r.PUT("/categories/:id", CategoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", CategoryHandler.DeleteCategory)

}
