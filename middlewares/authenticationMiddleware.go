package middlewares

import (
	"example.com/learn-golang/services"
	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.ValidateJWT(c)
		c.Next()
	}
}
