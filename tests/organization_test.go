package tests

import (
	"context"
	"testing"

	"campfire/internal/app/organization"
	"campfire/internal/database"
	"campfire/internal/repository"
	"campfire/pkg/utils"
)

func TestMain(m *testing.M) {
	utils.LoadEnv(".env")
	err := database.CreatePostgresConnection()
	if err != nil {
		panic(err)
	}
	defer database.ClosePostgresConnection()

	m.Run()
}

func TestCreateOrganization(t *testing.T) {
	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	userRepository := repository.NewUserRepositoryPostgres()
	s := organization.OrganizationService{
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
		_, _, err := s.CreateOrganization(context.TODO(), input)
		if err != nil {
			t.Fatalf("failed to create organization: %v", err)
		}
	})

	t.Run("Create organization with invalid data", func(t *testing.T) {
		invalidInput := organization.CreateOrganizationInput{
			OrganizationName: "",
			UserName:         "",
			Email:            "",
			Password:         "",
		}
		_, _, err := s.CreateOrganization(context.TODO(), invalidInput)

		if err == nil {
			t.Fatal("expected an error with invalid data, but got none")
		}
	})
}

func TestAddMemberToOrganization(t *testing.T) {
	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	userRepository := repository.NewUserRepositoryPostgres()
	s := organization.OrganizationService{
		UserRepository:         userRepository,
		OrganizationRepository: orgnRepository,
	}

	t.Run("Add member to organization successfully", func(t *testing.T) {
		input := organization.AddMemberInput{
			OrganizationId: 14,
			UserName:       "Pashmak",
			Email:          "pashmak@gmail.com",
			Password:       "pashmak",
		}

		_, err := s.AddMember(context.TODO(), input)
		if err != nil {
			t.Fatalf("failed to add member to organization: %v", err)
		}
	})
}
