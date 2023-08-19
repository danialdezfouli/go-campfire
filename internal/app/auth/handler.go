package auth

import (
	"campfire/internal/app/user"
	"campfire/pkg/config"
	"campfire/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService AuthService
	UserService user.UserService
}

func (h AuthHandler) Login(c *gin.Context) {
	input := LoginInput{}
	c.ShouldBindJSON(&input)
	user, err := h.AuthService.Attempt(c, input)

	if err != nil {
		c.JSON(err.GetCode(), err)
		return
	}

	secret := []byte(config.GetAccessTokenSecret())
	accessToken, err := h.AuthService.CreateAccessToken(c, user, secret)
	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":           user.Id,
			"name":         user.Name,
			"email":        user.Email,
			"access_token": accessToken,
			"expires_in":   1,
		},
	})
}

func (h AuthHandler) Me(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	claims, err := token.Parse(accessToken, config.GetAccessTokenSecret())
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.UserService.FindById(c, claims.User)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"claims": claims,
		"user":   user,
	})
}
