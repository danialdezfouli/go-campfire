package auth

import (
	"campfire/internal/domain"
	"campfire/pkg/exceptions"
	"campfire/pkg/token"
	"campfire/pkg/utils"
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	UserRepository domain.UserRepository
}

func NewAuthService(userRepository domain.UserRepository) AuthService {
	return AuthService{userRepository}
}

func (s AuthService) Attempt(ctx context.Context, input LoginInput) (*domain.User, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, &exceptions.ValidationError{
			Err: err,
		}
	}

	user, err := s.UserRepository.GetUserByEmail(ctx, input.Email, input.Subdomain)
	if err != nil {
		log.Printf("Error getting user: %s", err)
		return nil, &exceptions.InvalidLoginError{}
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return nil, &exceptions.InvalidLoginError{}
	}

	return user, nil
}

var signingKey = []byte("123")

func (s AuthService) CreateAccessToken(ctx context.Context, user *domain.User) (string, error) {
	claims := &token.Claims{
		UserID:         user.Id,
		OrganizationId: user.OrganizationId,
		Email:          user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	t, err := token.Generate(claims, signingKey)

	if err != nil {
		return "", err
	}

	return t, nil
}

func (s AuthService) VerifyToken(ctx context.Context, tokenStr string) (bool, error) {
	valid, err := token.Validate(tokenStr, signingKey)

	if err != nil {
		return false, err
	}

	if !valid {
		return false, errors.New("token is not valid")
	}

	return true, nil
}
