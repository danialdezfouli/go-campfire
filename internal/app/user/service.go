package user

import (
	"campfire/internal/domain"
	"campfire/pkg/exceptions"
	"context"
	"log"
)

type UserService struct {
	UserRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) UserService {
	return UserService{UserRepository: userRepository}
}

func (s UserService) FindById(ctx context.Context, id domain.UserId) (*domain.User, error) {
	user, err := s.UserRepository.GetUserById(ctx, id)

	if err != nil {
		log.Printf("failed to find user by id, %v", err)
		return nil, exceptions.InternalServerError
	}

	return user, nil

}
