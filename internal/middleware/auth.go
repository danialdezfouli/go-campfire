package middleware

import (
	"campfire/pkg/config"
	"campfire/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		signingKey := config.GetAccessTokenSecret()

		rawToken := c.GetHeader("Authorization")
		_, err := token.Parse(rawToken, signingKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Next()
	}
}
