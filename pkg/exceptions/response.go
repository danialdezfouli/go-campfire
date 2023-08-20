package exceptions

import "github.com/gin-gonic/gin"

func AbortWithError(c *gin.Context, err CustomError) {
	c.AbortWithStatusJSON(err.GetCode(), err)
}
