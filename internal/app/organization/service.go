package organization

import (
	"campfire/internal/domain"
	"context"

	"github.com/go-faker/faker/v4"
)

type OrganizationService struct {
	UserRepository         domain.UserRepository
	OrganizationRepository domain.OrganizationRepository
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{}
}

func generateRandomSubdomain() string {
	var sample struct {
		Subdomain string `faker:"username"`
	}
	faker.FakeData(&sample)
	return sample.Subdomain
}

func (s OrganizationService) CreateOrganization(ctx context.Context, input CreateOrganizationRequest) (*domain.Organization, *domain.User, error) {
	org := &domain.Organization{
		Name:      input.OrganizationName,
		Subdomain: generateRandomSubdomain(),
	}
	err := s.OrganizationRepository.CreateOrganization(ctx, org)

	if err != nil {
		return nil, nil, err
	}

	user := &domain.User{
		OrganizationId: org.Id,
		Name:           input.UserName,
		Email:          input.Email,
		Password:       input.Password,
		IsSuperAdmin:   true,
	}
	err = s.UserRepository.CreateUser(ctx, user)

	if err != nil {
		return org, nil, err
	}

	return org, user, nil
}
