package tests

import (
	"campfire/internal/app/organization"
	"campfire/internal/database"
	"campfire/internal/repository/postgresql"
	"campfire/pkg/utils"
	"context"
	"testing"
)

func TestCreateOrganization(t *testing.T) {
	utils.LoadEnv(".env")
	err := database.CreatePostgresConnection()
	defer database.ClosePostgresConnection()
	if err != nil {
		t.Fatal(err)
	}

	orgnRepository := postgresql.NewOrganizationRepositoryPostgres()
	userRepository := postgresql.NewUserRepositoryPostgres()

	s := organization.OrganizationService{UserRepository: userRepository, OrganizationRepository: orgnRepository}

	_, _, err = s.CreateOrganization(context.TODO(), organization.CreateOrganizationRequest{
		OrganizationName: "Test 2",
		UserName:         "Danial",
		Email:            "danial@gmail.com",
		Password:         "123456",
	})

	if err != nil {
		t.Fatal(err)
	}
}
