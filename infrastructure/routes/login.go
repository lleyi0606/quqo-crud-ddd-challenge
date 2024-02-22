package routes

import (
	"products-crud/infrastructure/controllers/handlers"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(r *gin.RouterGroup, p *base.Persistence) {

	AuthHandler := handlers.NewAuthorizationController(p)
	r.POST("/login", AuthHandler.Login)

}
