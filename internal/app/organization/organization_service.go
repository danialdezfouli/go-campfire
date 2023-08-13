package organization

import (
	"campfire/internal/domain"
	"campfire/internal/repository"
	"campfire/pkg/utils"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type OrganizationService struct {
	UserRepository         repository.UserRepositoryPostgres
	OrganizationRepository repository.OrganizationRepositoryPostgres
}

func (s OrganizationService) CreateOrganization(ctx context.Context, input CreateOrganizationInput) (*domain.Organization, *domain.User, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, nil, err
	}

	org := &domain.Organization{
		Name:      input.OrganizationName,
		Subdomain: input.Subdomain,
	}

	if err := s.OrganizationRepository.CreateOrganization(ctx, org); err != nil {
		return nil, nil, err
	}

	password, _ := utils.HashPassword(input.Password)
	user := &domain.User{
		OrganizationId: org.Id,
		Name:           input.UserName,
		Email:          input.Email,
		Password:       password,
		IsSuperAdmin:   true,
	}

	if err := s.UserRepository.CreateUser(ctx, user); err != nil {
		return org, nil, err
	}

	return org, user, nil
}

func (s OrganizationService) DeleteOrganization(ctx context.Context, id int) error {
	return s.OrganizationRepository.DeleteOrganization(ctx, id)
}

func (s OrganizationService) AddMember(ctx context.Context, input AddMemberInput) (*domain.User, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	org, err := s.OrganizationRepository.FindById(ctx, input.OrganizationId)

	if err != nil {
		return nil, err
	}

	if org == nil {
		return nil, fmt.Errorf("organization %d is not found", input.OrganizationId)
	}

	password, _ := utils.HashPassword(input.Password)
	user := &domain.User{
		OrganizationId: org.Id,
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
