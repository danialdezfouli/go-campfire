package organization

import (
	"campfire/internal/repository"

	"github.com/gin-gonic/gin"
)

func OrganizationRoutes(router *gin.Engine) {
	userRepository := repository.NewUserRepositoryPostgres()
	orgRepository := repository.NewOrganizationRepositoryPostgres()
	c := OrganizationHandler{OrganizationService: NewOrganizationService(orgRepository, userRepository)}

	router.POST("/api/v1/organizations/create", c.Create)
}
