package main

import (
	"log"

	"os"
	_ "products-crud/docs"
	base "products-crud/infrastructure/persistences"
	"products-crud/infrastructure/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"products-crud/infrastructure/config"

	"github.com/joho/godotenv"

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

	config.LoadConfiguration()

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Starting the application
	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = "8080" //localhost
	}
	log.Fatal(r.Run(":" + app_port))
}
