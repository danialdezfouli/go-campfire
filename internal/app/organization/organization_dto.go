package organization

type CreateOrganizationInput struct {
	UserName         string `validate:"required"`
	OrganizationName string `validate:"required"`
	Subdomain        string `validate:"required"`
	Email            string `validate:"required"`
	Password         string `validate:"required"`
}

type CreateOrganizationResponse struct {
}

type AddMemberInput struct {
	OrganizationId int    `validate:"required"`
	UserName       string `validate:"required"`
	Email          string `validate:"required"`
	Password       string `validate:"required"`
}
