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
		c.JSON(err.Code(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    user.Id,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
