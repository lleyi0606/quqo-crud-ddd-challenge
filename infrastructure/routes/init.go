package routes

import (
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func InitRouter(p *base.Persistence) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	Routes(r, p)

	return r
}

func Routes(r *gin.Engine, p *base.Persistence) {

	apiR := r.Group("")

	// List Injection
	ProductRoutes(apiR, p)
	InventoryRoutes(apiR, p)

}
