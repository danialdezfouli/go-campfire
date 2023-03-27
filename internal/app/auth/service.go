package auth

import (
	"campfire/internal/domain"
	"context"
	"log"
)

type AuthService struct {
	UserRepository domain.UserRepository
}

func NewAuthService(userRepository domain.UserRepository) AuthService {
	return AuthService{userRepository}
}

func (s AuthService) Login(ctx context.Context) (string, error) {
	user, err := s.UserRepository.GetUser(ctx, 1)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return "", err
	}
	return user.Name, nil
}

func (s AuthService) Register(ctx context.Context) (string, error) {
	user, err := s.UserRepository.GetUser(ctx, 1)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return "", err
	}
	return user.Name, nil
}

func (s AuthService) Logout(ctx context.Context) (string, error) {
	user, err := s.UserRepository.GetUser(ctx, 1)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return "", err
	}
	return user.Name, nil
}
