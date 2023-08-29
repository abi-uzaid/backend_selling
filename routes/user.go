package routes

import (
	"selling/handlers"
	"selling/pkg/mysql"
	"selling/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepository := repositories.NewUserRepository(mysql.DB)
	Handlers := handlers.NewUserHandler(userRepository)

	r.GET("/users", Handlers.FindUsers)
	r.GET("/users/:id", Handlers.GetUser)
	r.DELETE("/users/:id", Handlers.DeleteUser)
}