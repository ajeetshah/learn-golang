package routes

import (
	"example.com/learn-golang/constants"
	"example.com/learn-golang/controllers"
	"example.com/learn-golang/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/sign-up", controllers.Signup)
		publicRoutes.POST("/sign-in", controllers.Signin)
	}

	userRoutes := r.Group("/api")
	userRoutes.Use(middlewares.ValidateJWTAndRoleAndSetContext(constants.BasicRole))
	{
		userRoutes.GET("/books/:id", controllers.ReadBook)
		userRoutes.GET("/books", controllers.ReadBooks)
		userRoutes.POST("/books", controllers.CreateBook)
		userRoutes.PUT("/books/:id", controllers.UpdateBook)
		userRoutes.DELETE("/books/:id", controllers.DeleteBook)
	}

	adminRoutes := r.Group("/api/admin")
	adminRoutes.Use(middlewares.ValidateJWTAndRoleAndSetContext(constants.AdminRole))
	{
		adminRoutes.GET("/users", controllers.ReadUsers)
	}
}
