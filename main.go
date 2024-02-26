package main

import (
	"log"

	"os"
	_ "products-crud/docs"
	base "products-crud/infrastructure/persistences"
	"products-crud/infrastructure/routes"

	"products-crud/infrastructure/config"

	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

// func initTracer() (*sdktrace.TracerProvider, error) {
// 	exporter, err := stdout.New(stdout.WithPrettyPrint())
// 	if err != nil {
// 		return nil, err
// 	}
// 	tp := sdktrace.NewTracerProvider(
// 		sdktrace.WithSampler(sdktrace.AlwaysSample()),
// 		sdktrace.WithBatcher(exporter),
// 	)
// 	otel.SetTracerProvider(tp)
// 	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
// 	return tp, nil
// }

// func getUser(c *gin.Context, id string) string {
// 	// Pass the built-in `context.Context` object from http.Request to OpenTelemetry APIs
// 	// where required. It is available from gin.Context.Request.Context()
// 	_, span := tracer.Start(c.Request.Context(), "getUser", oteltrace.WithAttributes(attribute.String("id", id)))
// 	defer span.End()
// 	if id == "123" {
// 		return "otelgin tester"
// 	}
// 	return "unknown"
// }
