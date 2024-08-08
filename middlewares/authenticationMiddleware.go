package middlewares

import (
	"example.com/learn-golang/utils"
	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.ValidateJWT(c)
		c.Next()
	}
}
