package middlewares

import (
	"example.com/learn-golang/models"
	"example.com/learn-golang/utils"
	"github.com/gin-gonic/gin"
)

func ValidateJWTAndRoleAndSetContext(eligibleRole models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.ValidateJWTAndRoleAndSetContext(c, eligibleRole)
		c.Next()
	}
}
