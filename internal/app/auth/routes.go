package auth

import (
	"campfire/internal/middleware"
	"campfire/internal/repository"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	userRepository := repository.NewUserRepositoryPostgres()
	c := AuthHandler{
		AuthService: NewAuthService(userRepository),
	}

	public := router.Group("/api/v1/auth")
	private := router.Group("/api/v1/auth")

	public.POST("login", c.Login)

	private.Use(middleware.AuthMiddleware())
	private.GET("me", c.Me)
}
