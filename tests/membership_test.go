package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"campfire/internal/app/membership"
	"campfire/internal/database"
	"campfire/internal/repository"
	"campfire/pkg/config"
)

func TestAddMemberToOrganization(t *testing.T) {
	config.LoadEnv(".env")
	err := database.CreatePostgresConnection()
	assert.NoError(t, err)
	defer database.ClosePostgresConnection()

	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	userRepository := repository.NewUserRepositoryPostgres()
	s := membership.MembershipService{
		OrganizationRepository: orgnRepository,
		UserRepository:         userRepository,
	}

	t.Run("Add member to organization successfully", func(t *testing.T) {
		input := membership.AddMemberInput{
			OrganizationId: 14,
			UserName:       "Pashmak",
			Email:          "pashmak@gmail.com",
			Password:       "pashmak",
		}

		_, err := s.AddMember(context.Background(), input)
		assert.NoError(t, err)
	})
}
