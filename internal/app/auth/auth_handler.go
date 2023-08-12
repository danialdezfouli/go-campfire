package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService AuthService
}

func (h AuthHandler) Login(c *gin.Context) {
	data, err := h.AuthService.Login(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func (h AuthHandler) Register(c *gin.Context) {
	data, err := h.AuthService.Register(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": data,
	})
}

func (h AuthHandler) Logout(c *gin.Context) {
	data, err := h.AuthService.Logout(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
