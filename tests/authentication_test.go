package tests

import (
	"campfire/internal/app/auth"
	"campfire/internal/app/organization"
	"campfire/internal/database"
	"campfire/internal/domain"
	"campfire/internal/repository"
	"campfire/pkg/config"
	"campfire/pkg/token"
	"campfire/pkg/utils"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func setupTestEnvironment(t *testing.T) {
	config.LoadEnv(".env")
	assert.NoError(t, database.CreatePostgresConnection())
}

func tearDownTestEnvironment() {
	database.ClosePostgresConnection()
}

func createTestOrganization(t *testing.T) (*domain.Organization, *domain.User) {
	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	userRepository := repository.NewUserRepositoryPostgres()
	orgService := organization.OrganizationService{
		UserRepository:         userRepository,
		OrganizationRepository: orgnRepository,
	}

	org, user, err := orgService.CreateOrganization(context.Background(), organization.CreateOrganizationInput{
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

func deleteTestOrganization(t *testing.T, id domain.OrganizationId) {
	userRepository := repository.NewUserRepositoryPostgres()
	orgnRepository := repository.NewOrganizationRepositoryPostgres()
	orgService := organization.OrganizationService{
		UserRepository:         userRepository,
		OrganizationRepository: orgnRepository,
	}

	err := orgService.DeleteOrganization(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserLogin(t *testing.T) {
	setupTestEnvironment(t)
	defer tearDownTestEnvironment()

	userRepository := repository.NewUserRepositoryPostgres()
	authService := auth.AuthService{
		UserRepository: userRepository,
	}

	t.Run("User can login successfully", func(t *testing.T) {
		org, createdUser := createTestOrganization(t)
		defer deleteTestOrganization(t, org.Id)

		_, err := authService.Attempt(context.Background(), auth.LoginInput{
			Subdomain: org.Subdomain,
			Email:     createdUser.Email,
			Password:  "secret",
		})

		assert.NoError(t, err, "expected no error when logging in")
	})

	t.Run("User does not login with invalid email", func(t *testing.T) {
		org, _ := createTestOrganization(t)
		defer deleteTestOrganization(t, org.Id)

		_, err := authService.Attempt(context.Background(), auth.LoginInput{
			Subdomain: org.Subdomain,
			Email:     "different-email@gmail.com",
			Password:  "secret",
		})

		assert.Error(t, err, "expected error when logging in with invalid email")
	})

}

func TestCreateToken(t *testing.T) {
	setupTestEnvironment(t)
	defer tearDownTestEnvironment()

	signingKey := []byte("123")
	userRepository := repository.NewUserRepositoryPostgres()
	authService := auth.AuthService{
		UserRepository: userRepository,
	}

	t.Run("User get a new token", func(t *testing.T) {
		org, user := createTestOrganization(t)
		defer deleteTestOrganization(t, org.Id)

		token, err := authService.CreateAccessToken(context.Background(), user, signingKey)
		assert.NoError(t, err, "expected no error when creating token")
		assert.NotEmpty(t, token, "expected token to be non-empty")

		log.Printf("generated token: %v", token)
	})

	t.Run("Given token is valid", func(t *testing.T) {
		claims := &token.Claims{
			User:         1,
			Organization: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			},
		}

		created, err := token.Generate(claims, signingKey)
		assert.NoError(t, err, "expected no error when generating token")
		assert.NotEmpty(t, created, "expected created token to be non-empty")

		fmt.Println(created)

		_, err = token.Validate(created, signingKey)
		assert.NoError(t, err, "expected no error when validating token")
	})

	t.Run("Token validation fails on expired dates", func(t *testing.T) {
		claims := &token.Claims{
			User:         1,
			Organization: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
			},
		}

		created, _ := token.Generate(claims, signingKey)
		_, err := token.Validate(created, signingKey)
		assert.Error(t, err, "token validation doesn't work right")
	})

	t.Run("Token validation fails on wrong signing key", func(t *testing.T) {
		claims := &token.Claims{
			User:         1,
			Organization: 1,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			},
		}

		created, _ := token.Generate(claims, []byte("xyz"))
		_, err := token.Validate(created, signingKey)

		assert.Error(t, err, "expected error when validating expired token")

	})
}
