package routes

import (
	"log"
	"products-crud/infrastructure/controllers/handlers"
	"products-crud/infrastructure/controllers/middleware"
	base "products-crud/infrastructure/persistences"

	"github.com/gin-gonic/gin"
	"github.com/honeycombio/honeycomb-opentelemetry-go"
	"github.com/honeycombio/otel-config-go/otelconfig"
	// "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func InitRouter(p *base.Persistence) *gin.Engine {

	// Enable multi-span attributes
	bsp := honeycomb.NewBaggageSpanProcessor()

	// Use the Honeycomb distro to set up the OpenTelemetry SDK
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithSpanProcessor(bsp),
	)
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}
	defer otelShutdown()

	// Initialize HTTP handler instrumentation
	// handler := http.HandlerFunc(httpHandler)
	// wrappedHandler := otelhttp.NewHandler(handler, "hello")
	// http.Handle("/hello", wrappedHandler)

	r := gin.New()

	r.Use(gin.Recovery())

	Routes(r, p)

	return r
}

// func WrapHandler(handler gin.HandlerFunc, routeName string) gin.HandlerFunc {
// 	wrappedHandler := otelhttp.NewHandler(http.HandlerFunc(handler), routeName)
// 	return func(c *gin.Context) {
// 		wrappedHandler.ServeHTTP(c.Writer, c.Request)
// 	}

// 	otelgin.Middleware()
// }

func Routes(r *gin.Engine, p *base.Persistence) {

	apiR := r.Group("")

	// Public routes with only honeycomb middleware
	apiR.Use(middleware.HoneycombHandler())
	LoginRoutes(apiR, p)

	// Protected routes with middleware
	AuthHandler := handlers.NewAuthorizationController(p)
	protectedRoutes := apiR.Group("")
	protectedRoutes.Use(middleware.AuthHandler(p), middleware.HoneycombHandler())
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
