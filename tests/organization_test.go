package tests

import (
	"campfire/internal/app/organization"
	"campfire/internal/database"
	"campfire/internal/repository"
	"campfire/pkg/utils"
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
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

func generateRandomSubdomain() string {
	var sample struct {
		Subdomain string `faker:"username"`
	}
	faker.FakeData(&sample)
	return sample.Subdomain
}

func TestCreateOrganization(t *testing.T) {
	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	userRepository := repository.NewUserRepositoryPostgres()
	s := organization.OrganizationService{
		UserRepository:         userRepository,
		OrganizationRepository: orgnRepository,
	}

	t.Run("Create organization successfully", func(t *testing.T) {
		_, _, err := s.CreateOrganization(context.TODO(), organization.CreateOrganizationRequest{
			UserName:         "Danial",
			OrganizationName: "Example Organization",
			Subdomain:        generateRandomSubdomain(),
			Email:            "danial@gmail.com",
			Password:         "secret",
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create organization with invalid data", func(t *testing.T) {
		_, _, err := s.CreateOrganization(context.TODO(), organization.CreateOrganizationRequest{
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
		_, err := s.AddMember(context.TODO(), organization.AddMemberRequest{
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
