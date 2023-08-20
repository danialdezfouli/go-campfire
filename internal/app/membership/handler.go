package membership

import (
	"campfire/pkg/exceptions"
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
		exceptions.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
