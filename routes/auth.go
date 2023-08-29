package routes

import (
	"selling/handlers"
	"selling/pkg/middleware"
	"selling/pkg/mysql"
	"selling/repositories"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	userRepository := repositories.NewAuthRepository(mysql.DB)
	Handlers := handlers.HandlerAuth(userRepository)

	r.POST("/register", Handlers.Register)
	r.POST("/login", Handlers.Login)
	r.GET("/check-auth", middleware.Auth(Handlers.CheckAuth))
}
