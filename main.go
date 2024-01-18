package main

import (
	"log"

	"products-crud/infrastructure/controllers/handlers"
	"products-crud/infrastructure/routes"

	"os"
	base "products-crud/infrastructure/persistences"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	p, err := base.NewPersistence()
	if err != nil {
		panic(err)
	}

	// Migrations
	p.Automigrate()

	// Defer close
	defer p.Close()

	// Routes
	r := routes.InitRouter(p)

	r := gin.Default()

	products := handlers.NewProducts(services.Product)
	//product routes
	r.POST("/products", products.AddProduct)
	r.GET("/products", products.GetProducts)
	r.GET("/products/:id", products.GetProduct)
	r.PUT("/products/:id", products.UpdateProduct)
	r.DELETE("/products/:id", products.DeleteProduct)
	r.GET("/products/search", products.SearchProducts)

	//Starting the application
	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = "8080" //localhost
	}
	log.Fatal(r.Run(":" + app_port))
}
