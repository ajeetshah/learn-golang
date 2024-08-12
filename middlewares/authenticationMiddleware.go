package middlewares

import (
	"example.com/learn-golang/utils"
	"github.com/gin-gonic/gin"
)

func ValidateTokenAndSetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.ValidateJWTAndSetRole(c)
		c.Next()
	}
}
