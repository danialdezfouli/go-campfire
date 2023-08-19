package membership

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MembershipHandler struct {
	MembershipService MembershipService
}

func (h MembershipHandler) AddMember(c *gin.Context) {
	input := AddMemberInput{}
	c.ShouldBindJSON(&input)

	user, err := h.MembershipService.AddMember(c, input)

	if err != nil {
		c.AbortWithStatusJSON(err.GetCode(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
