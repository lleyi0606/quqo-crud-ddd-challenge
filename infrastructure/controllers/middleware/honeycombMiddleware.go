package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/honeycombio/honeycomb-opentelemetry-go"
	"github.com/honeycombio/otel-config-go/otelconfig"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// Custom middleware for OpenTelemetry instrumentation
func HoneycombHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Enable multi-span attributes
		bsp := honeycomb.NewBaggageSpanProcessor()
		// Use the Honeycomb distro to set up the OpenTelemetry SDK
		otelShutdown, err := otelconfig.ConfigureOpenTelemetry(
			otelconfig.WithSpanProcessor(bsp),
		)
		if err != nil {
			log.Fatalf("error setting up OTel SDK - %s", err)
		}
		defer otelShutdown()

		otelgin.Middleware("gin-server")(c)

		c.Next()
	}
}
