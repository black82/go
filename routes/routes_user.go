package routes

import (
	"awesomeProject/config"
	"awesomeProject/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users").Use(config.IsAuthenticated())
	{
		userRoutes.GET("/", handlers.GetUsers)
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.GET("/:id", handlers.GetUser)
		userRoutes.PUT("/:id", handlers.UpdateUser)
		userRoutes.DELETE("/:id", handlers.DeleteUser)
	}
	r.POST("/login", handlers.Login)
	r.POST("/sign-up", handlers.Signup)
}
