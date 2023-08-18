package organization

import (
	"campfire/internal/domain"
	"campfire/pkg/exceptions"
	"campfire/pkg/utils"
	"context"
)

func (s OrganizationService) AddMember(ctx context.Context, input AddMemberInput) (*domain.User, error) {
	if err := validateInput(input); err != nil {
		return nil, err
	}

	organization, err := s.OrganizationRepository.FindById(ctx, input.OrganizationId)
	if err != nil {
		return nil, err
	}

	if organization == nil {
		return nil, exceptions.OrganizationNotFound(input.OrganizationId)
	}

	password, _ := utils.HashPassword(input.Password)
	user := &domain.User{
		OrganizationId: organization.Id,
		Name:           input.UserName,
		Email:          input.Email,
		Password:       password,
		IsSuperAdmin:   false,
	}

	if err := s.UserRepository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
