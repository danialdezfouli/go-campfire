package tests

import (
	"campfire/internal/app/auth"
	"campfire/internal/app/organization"
	"campfire/internal/domain"
	"campfire/internal/repository"
	"campfire/pkg/token"
	"campfire/pkg/utils"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createOrganization(t *testing.T) (*domain.Organization, *domain.User) {
	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	userRepository := repository.NewUserRepositoryPostgres()
	orgService := organization.OrganizationService{
		UserRepository:         userRepository,
		OrganizationRepository: orgnRepository,
	}

	org, user, err := orgService.CreateOrganization(context.TODO(), organization.CreateOrganizationInput{
		UserName:         "testing-auth",
		OrganizationName: "testing-auth",
		Subdomain:        utils.GenerateRandomSubdomain(),
		Email:            "testing-auth@gmail.com",
		Password:         "secret",
	})

	if err != nil {
		t.Fatal(err)
	}

	return org, user
}

func deleteOrganization(id int) {
	userRepository := repository.NewUserRepositoryPostgres()
	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	orgService := organization.OrganizationService{
		UserRepository:         userRepository,
		OrganizationRepository: orgnRepository,
	}
	defer orgService.DeleteOrganization(context.TODO(), id)
}

func TestUserLogin(t *testing.T) {
	userRepository := repository.NewUserRepositoryPostgres()
	authService := auth.AuthService{
		UserRepository: userRepository,
	}

	t.Run("User can login successfully", func(t *testing.T) {
		org, createdUser := createOrganization(t)
		defer deleteOrganization(org.Id)

		_, err := authService.Attempt(context.TODO(), auth.LoginInput{
			Email:    createdUser.Email,
			Password: createdUser.Password,
		})

		if err != nil {
			t.Fatal("failed to login", err)
		}

	})

	t.Run("User does not login with invalid email", func(t *testing.T) {
		org, createdUser := createOrganization(t)
		defer deleteOrganization(org.Id)

		_, err := authService.Attempt(context.TODO(), auth.LoginInput{
			Email:    "different-email@gmail.com",
			Password: createdUser.Password,
		})

		if err == nil {
			t.Fatal("it should not login")
		}
	})

}

func TestCreateToken(t *testing.T) {
	signingKey := []byte("123")
	userRepository := repository.NewUserRepositoryPostgres()
	authService := auth.AuthService{
		UserRepository: userRepository,
	}

	t.Run("user get a new token", func(t *testing.T) {
		org, user := createOrganization(t)
		defer deleteOrganization(org.Id)

		token, err := authService.CreateAccessToken(context.TODO(), user)
		if err != nil {
			t.Fatal(err)
		}

		if token == "" {
			t.Fatal("token is empty")
		}

		log.Printf("generated token: %v", token)
	})

	t.Run("given token is valid", func(t *testing.T) {
		claims := &token.Claims{
			UserID:         1,
			OrganizationId: 1,
			Email:          "testing-token@gmail.com",
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			},
		}

		created, err := token.Generate(claims, signingKey)
		if err != nil {
			t.Fatal(err)
		}

		if created == "" {
			t.Fatal("token is empty")
		}

		fmt.Println(created)

		verified, err := token.Validate(created, signingKey)

		if err != nil {
			t.Fatal(err)
		}

		if !verified {
			t.Fatal("token is not verified")
		}

	})

	t.Run("token validation failes on expired dates", func(t *testing.T) {
		claims := &token.Claims{
			UserID:         1,
			OrganizationId: 1,
			Email:          "testing-token@gmail.com",
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
			},
		}

		created, _ := token.Generate(claims, signingKey)
		_, err := token.Validate(created, signingKey)

		if err == nil {
			t.Fatal("token validation doesn't work right")
		}
	})

	t.Run("token validation failes on wrong signing key", func(t *testing.T) {
		claims := &token.Claims{
			UserID:         1,
			OrganizationId: 1,
			Email:          "testing-token@gmail.com",
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			},
		}

		created, _ := token.Generate(claims, []byte("xyz"))
		valid, _ := token.Validate(created, signingKey)

		if valid {
			t.Fatal("token validation should not pass wrong signingkey")
		}
	})
}
