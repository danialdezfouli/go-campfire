package auth

import (
	"campfire/pkg/exceptions"
	"campfire/pkg/token"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) (*token.Claims, exceptions.CustomError) {
	claims, exists := c.Get("user")
	if !exists {
		return nil, exceptions.Unauthenticated
	}

	userClaims := claims.(*token.Claims)

	return userClaims, nil
}
