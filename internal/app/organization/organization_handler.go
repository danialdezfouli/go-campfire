package organization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	OrganizationService OrganizationService
}

func (h OrganizationHandler) Create(c *gin.Context) {
	org, user, err := h.OrganizationService.CreateOrganization(c, CreateOrganizationRequest{
		UserName:         "",
		OrganizationName: "",
		Email:            "",
		Password:         "",
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"organization": org,
		"user":         user,
	})
}
