package auth

import (
	"campfire/internal/repository/postgresql"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	userRepositoryDb := postgresql.NewUserRepositoryPostgres()
	c := AuthHandler{AuthService: NewAuthService(userRepositoryDb)}

	route.POST("/login", c.Login)
	route.POST("/register", c.Register)
	route.DELETE("/logout", c.Logout)
}
