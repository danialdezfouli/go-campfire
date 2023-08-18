package organization

import (
	"campfire/internal/domain"
	"campfire/pkg/utils"
	"context"
)

func (s OrganizationService) CreateOrganization(ctx context.Context, input CreateOrganizationInput) (*domain.Organization, *domain.User, error) {
	if err := validateInput(input); err != nil {
		return nil, nil, err
	}

	org, err := s.createOrganization(ctx, input)
	if err != nil {
		return nil, nil, err
	}

	user, err := s.createFirstUserForOrganization(ctx, org.Id, input)
	if err != nil {
		return org, nil, err
	}

	return org, user, nil
}

func (s OrganizationService) createOrganization(ctx context.Context, input CreateOrganizationInput) (*domain.Organization, error) {
	org := &domain.Organization{
		Name:      input.OrganizationName,
		Subdomain: input.Subdomain,
	}

	if err := s.OrganizationRepository.CreateOrganization(ctx, org); err != nil {
		return nil, err
	}

	return org, nil
}

func (s OrganizationService) createFirstUserForOrganization(ctx context.Context, orgID int, input CreateOrganizationInput) (*domain.User, error) {
	password, _ := utils.HashPassword(input.Password)
	user := &domain.User{
		OrganizationId: orgID,
		Name:           input.UserName,
		Email:          input.Email,
		Password:       password,
		IsSuperAdmin:   true,
	}

	if err := s.UserRepository.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
