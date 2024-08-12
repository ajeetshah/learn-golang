package routes

import (
	"example.com/learn-golang/controllers"
	"example.com/learn-golang/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/sign-up", controllers.Signup)
		publicRoutes.POST("/sign-in", controllers.Signin)
	}

	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middlewares.ValidateTokenAndSetContext())
	{
		protectedRoutes.GET("/books/:id", controllers.ReadBook)
		protectedRoutes.GET("/books", controllers.ReadBooks)
		protectedRoutes.POST("/books", controllers.CreateBook)
		protectedRoutes.PUT("/books/:id", controllers.UpdateBook)
		protectedRoutes.DELETE("/books/:id", controllers.DeleteBook)
	}
}
