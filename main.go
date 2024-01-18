package main

import (
	"log"

	"products-crud/infrastructure/controllers/handlers"

	"os"
	"products-crud/infrastructure/persistence"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	services, err := persistence.NewRepositories(os.Getenv("DB_PASSWORD"))
	if err != nil {
		panic(err)
	}
	services.Automigrate()

	products := handlers.NewProducts(services.Product)

	r := gin.Default()

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
