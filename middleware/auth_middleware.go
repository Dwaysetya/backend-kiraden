package middlewares

import (
	"backend/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware()gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")


		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is requaired",
			})
			c.Abort()
			return
		}

		tokenString =  strings.TrimPrefix(tokenString, "Bearer ")
		claims := &jwt.RegisteredClaims{}

		token, err :=  jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}
		c.Set("name", claims.Subject)

		c.Next()
	}
}