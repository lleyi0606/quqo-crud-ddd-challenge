package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	"products-crud/infrastructure/controllers/middleware"
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

	// Public routes without middleware
	LoginRoutes(apiR, p)

	// Protected routes with middleware
	AuthHandler := handlers.NewAuthorizationController(p)
	protectedRoutes := apiR.Group("")
	protectedRoutes.Use(middleware.AuthHandler(p))
	{
		ProductRoutes(protectedRoutes, p)
		InventoryRoutes(protectedRoutes, p)
		ImageRoutes(protectedRoutes, p)
		CategoryRoutes(protectedRoutes, p)
		CustomerRoutes(protectedRoutes, p)
		OrderRoutes(protectedRoutes, p)
		OrderedItemRoutes(protectedRoutes, p)
		protectedRoutes.POST("/logout", AuthHandler.Logout)
	}

}
