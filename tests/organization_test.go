package tests

import (
	"context"
	"testing"

	"campfire/internal/app/organization"
	"campfire/internal/database"
	"campfire/internal/repository"
	"campfire/pkg/config"
	"campfire/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrganization(t *testing.T) {
	// database connection
	config.LoadEnv(".env")
	assert.NoError(t, database.CreatePostgresConnection())
	defer database.ClosePostgresConnection()

	// initilizing
	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	userRepository := repository.NewUserRepositoryPostgres()
	orgService := organization.OrganizationService{
		UserRepository:         userRepository,
		OrganizationRepository: orgnRepository,
	}

	t.Run("Create organization successfully", func(t *testing.T) {
		input := organization.CreateOrganizationInput{
			UserName:         "Danial",
			OrganizationName: "Example Organization",
			Subdomain:        utils.GenerateRandomSubdomain(),
			Email:            "danial@gmail.com",
			Password:         "secret",
		}
		_, _, err := orgService.CreateOrganization(context.Background(), input)
		assert.NoError(t, err)
		// assert.NoError(t, err, "expected no error when creating organization")

	})

	t.Run("Create organization with invalid data", func(t *testing.T) {
		invalidInput := organization.CreateOrganizationInput{
			OrganizationName: "",
			UserName:         "",
			Email:            "",
			Password:         "",
		}
		_, _, err := orgService.CreateOrganization(context.Background(), invalidInput)
		assert.Error(t, err, "expected error with invalid data when creating organization")

	})
}
