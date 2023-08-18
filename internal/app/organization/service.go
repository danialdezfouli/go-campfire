package organization

import (
	"campfire/internal/domain"
	"context"

	"github.com/go-playground/validator/v10"
)

type OrganizationService struct {
	UserRepository         domain.UserRepository
	OrganizationRepository domain.OrganizationRepository
}

func validateInput(input any) error {
	validate := validator.New()
	return validate.Struct(input)
}

func (s OrganizationService) DeleteOrganization(ctx context.Context, id int) error {
	return s.OrganizationRepository.DeleteOrganization(ctx, id)
}
