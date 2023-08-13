package tests

import (
	"campfire/internal/app/organization"
	"campfire/internal/database"
	"campfire/internal/repository"
	"campfire/pkg/utils"
	"context"
	"testing"
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
		_, _, err := s.CreateOrganization(context.TODO(), organization.CreateOrganizationInput{
			UserName:         "Danial",
			OrganizationName: "Example Organization",
			Subdomain:        utils.GenerateRandomSubdomain(),
			Email:            "danial@gmail.com",
			Password:         "secret",
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create organization with invalid data", func(t *testing.T) {
		_, _, err := s.CreateOrganization(context.TODO(), organization.CreateOrganizationInput{
			OrganizationName: "",
			UserName:         "",
			Email:            "",
			Password:         "",
		})

		if err == nil {
			t.Fatal("it should fail with invalid data.")
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
		_, err := s.AddMember(context.TODO(), organization.AddMemberInput{
			OrganizationId: 14,
			UserName:       "Pashmak",
			Email:          "pashmak@gmail.com",
			Password:       "pashmak",
		})

		if err != nil {
			t.Fatal(err)
		}
	})

}
