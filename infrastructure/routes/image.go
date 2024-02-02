package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func ImageRoutes(r *gin.RouterGroup, p *base.Persistence) {

	ImageHandler := handlers.NewImageController(p)

	r.POST("/images", ImageHandler.AddImage)
	r.GET("/images/:id", ImageHandler.GetImage)
	r.DELETE("/images/:id", ImageHandler.DeleteImage)

}
