package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func ImageRoutes(r *gin.RouterGroup, p *base.Persistence) {

	ImageHandler := handlers.NewImageController(p)

	r.POST("/images", ImageHandler.AddImage)
	// r.POST("/images/all", inventoryHandler.AddInventories)
	r.GET("/images/:id", ImageHandler.GetImage)
	// r.PUT("/images/:id", inventoryHandler.UpdateStock)
	// r.PUT("/images/:id", inventoryHandler.UpdateInventory)
	r.DELETE("/images/:id", ImageHandler.DeleteImage)
	// r.GET("/images/search", inventoryHandler.SearchInventory)

}
