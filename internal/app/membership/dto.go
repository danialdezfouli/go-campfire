package membership

import "campfire/internal/domain"

type AddMemberInput struct {
	OrganizationId domain.OrganizationId `validate:"required"`
	UserName       string                `validate:"required"`
	Email          string                `validate:"required"`
	Password       string                `validate:"required"`
}
