package user

import (
	"campfire/internal/domain"
)

type UserService struct {
	UserRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) UserService {
	return UserService{UserRepository: userRepository}
}
