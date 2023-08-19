package auth

import (
	"campfire/internal/domain"
	"campfire/pkg/config"
	"campfire/pkg/exceptions"
	"campfire/pkg/token"
	"campfire/pkg/utils"
	"campfire/pkg/validations"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	UserRepository domain.UserRepository
}

func NewAuthService(userRepository domain.UserRepository) AuthService {
	return AuthService{UserRepository: userRepository}
}

func (s AuthService) Attempt(ctx context.Context, input LoginInput) (*domain.User, exceptions.CustomError) {
	if err := validations.Validate(input); err != nil {
		return nil, exceptions.NewValidationError(err)
	}

	user, err := s.UserRepository.GetUserByEmail(ctx, input.Email, input.Subdomain)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, exceptions.InvalidLogin
	}

	if user == nil {
		return nil, exceptions.InvalidLogin
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return nil, exceptions.InvalidLogin
	}

	return user, nil
}

func (s AuthService) CreateAccessToken(ctx context.Context, user *domain.User, signingKey []byte) (string, exceptions.CustomError) {
	claims := &token.Claims{
		User:         user.Id,
		Organization: user.OrganizationId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	str, err := token.Generate(claims, signingKey)

	if err != nil {
		log.Printf("failed to generate token: %v", err)
		return "", exceptions.InternalServerError
	}

	return str, nil
}

func (s AuthService) VerifyToken(ctx context.Context, tokenString string, signingKey []byte) (bool, exceptions.CustomError) {
	_, err := token.Validate(tokenString, signingKey)
	if err != nil {
		log.Printf("invalid token: %v", err)
		return false, exceptions.InvalidToken
	}

	return true, nil
}

func (s AuthService) ParseToken(ctx *gin.Context) (*token.Claims, exceptions.CustomError) {
	accessToken := ctx.GetHeader("Authorization")
	claims, err := token.Parse(accessToken, config.GetAccessTokenSecret())
	if err != nil {
		return nil, &exceptions.RequestError{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}
	}

	return claims, nil
}
