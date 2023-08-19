package membership

import (
	"campfire/internal/domain"
	"campfire/pkg/exceptions"
	"campfire/pkg/utils"
	"campfire/pkg/validations"
	"context"
)

type MembershipService struct {
	OrganizationRepository domain.OrganizationRepository
	UserRepository         domain.UserRepository
}

func NewMembershipService(OrganizationRepository domain.OrganizationRepository, userRepository domain.UserRepository) MembershipService {
	return MembershipService{
		OrganizationRepository,
		userRepository,
	}
}

func (s MembershipService) AddMember(ctx context.Context, input AddMemberInput) (*domain.User, exceptions.CustomError) {
	if err := validations.Validate(input); err != nil {
		return nil, exceptions.NewValidationError(err)
	}

	organization, err := s.OrganizationRepository.FindById(ctx, input.OrganizationId)
	if err != nil {
		return nil, exceptions.InternalServerError("failed to fetch organization", err)
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
		return nil, exceptions.InternalServerError("failed to create user", err)
	}

	return user, nil
}
