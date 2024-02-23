package middleware

import (
	"log"
	"net/http"
	"os"
	"products-crud/domain/entity"
	"products-crud/infrastructure/implementations/authorization"
	base "products-crud/infrastructure/persistences"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthHandler(p *base.Persistence) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		log.Print("token in middleware: ", token)
		b := "Bearer "
		if !strings.Contains(token, b) {
			c.JSON(http.StatusForbidden, gin.H{"message": "Your request is not authorized", "status": entity.StatusError, "data": nil})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(http.StatusForbidden, gin.H{"message": "An authorization token was not supplied", "status": entity.StatusError, "data": nil})
			c.Abort()
			return
		}

		auth := authorization.NewAuthorizationRepository(p)
		key := os.Getenv("DATABASE_URL")
		v, err := auth.ValidateToken(t[1], key)

		// keyJWT := os.Getenv("KeyAuth")
		// v2, err2 := authorization.ValidateToken(t[1], keyJWT)

		if err != nil || v != nil && !v.Valid {
			c.JSON(http.StatusForbidden, gin.H{"message": "Invalid authorization token", "status": entity.StatusError, "data": nil})
			c.Abort()
			return
		}

		//catch token
		var tokenCatches jwt.Token
		if v.Valid {
			tokenCatches = *v
		}

		userIDInterface := tokenCatches.Claims.(jwt.MapClaims)["user_id"]
		if userID, ok := userIDInterface.(string); ok {
			// Now userID is of type string
			log.Println("!!! USER ID IS: ", userID)
			if userID == "" || userID == "0" {
				c.JSON(http.StatusForbidden, gin.H{"message": "Invalid authorization token", "status": entity.StatusError, "data": nil})
				c.Abort()
				return
			}
			c.Set("userID", userID)
		}

		c.Next()

	}
}
