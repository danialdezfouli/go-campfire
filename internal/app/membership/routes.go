package membership

import (
	"campfire/internal/repository"

	"github.com/gin-gonic/gin"
)

func MembershipRoutes(router *gin.Engine) {
	userRepository := repository.NewUserRepositoryPostgres()
	orgRepository := repository.NewOrganizationRepositoryPostgres()
	c := MembershipHandler{MembershipService: NewMembershipService(orgRepository, userRepository)}

	router.POST("/api/v1/organizations/:id/members", c.AddMember)
}
