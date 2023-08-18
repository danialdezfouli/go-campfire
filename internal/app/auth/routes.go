package auth

import (
	"campfire/internal/repository"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	userRepositoryDb := repository.NewUserRepositoryPostgres()
	c := AuthHandler{AuthService: NewAuthService(userRepositoryDb)}

	router.POST("/api/v1/login", c.Login)
}
