package auth

import (
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
		c.JSON(err.GetCode(), err)
		return
	}

	secret := []byte("123")
	accessToken, err := h.AuthService.CreateAccessToken(c, user, secret)
	if err != nil {
		c.JSON(err.GetCode(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":            user.Id,
			"name":          user.Name,
			"email":         user.Email,
			"access_token":  accessToken,
			"refresh_token": "refresh-token",
			"expires_in":    3600,
		},
	})
}
