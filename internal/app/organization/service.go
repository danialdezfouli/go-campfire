package organization

import (
	"campfire/internal/domain"
	"context"
)

type OrganizationService struct {
	UserRepository         domain.UserRepository
	OrganizationRepository domain.OrganizationRepository
}

func (s OrganizationService) DeleteOrganization(ctx context.Context, id int) error {
	return s.OrganizationRepository.DeleteOrganization(ctx, id)
}
