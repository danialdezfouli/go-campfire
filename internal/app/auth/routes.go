package auth

import (
	"campfire/internal/repository"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	userRepository := repository.NewUserRepositoryPostgres()
	c := AuthHandler{AuthService: NewAuthService(userRepository)}

	router.POST("/api/v1/auth/login", c.Login)
}
