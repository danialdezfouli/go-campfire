package auth

import (
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
		switch err.(type) {
		case nil:
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
		case *exceptions.InvalidLoginError:
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
		case *exceptions.ValidationError:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

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
