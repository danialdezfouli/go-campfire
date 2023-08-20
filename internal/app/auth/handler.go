package auth

import (
	"campfire/pkg/config"
	"campfire/pkg/exceptions"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService AuthService
}

func (h AuthHandler) Login(c *gin.Context) {
	input := LoginInput{}
	c.ShouldBindJSON(&input)
	user, err := h.AuthService.Attempt(c, input)

	if err != nil {
		exceptions.AbortWithError(c, err)
		return
	}

	secret := []byte(config.GetAccessTokenSecret())
	accessToken, err := h.AuthService.CreateAccessToken(c, user, secret)
	if err != nil {
		exceptions.AbortWithError(c, err)
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
	claims, err := GetUser(c)
	if err != nil {
		exceptions.AbortWithError(c, err)
		return
	}

	user, err := h.AuthService.FindUserById(c, claims.User)
	if err != nil {
		exceptions.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
