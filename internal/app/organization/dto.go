package organization

type CreateOrganizationRequest struct {
	UserName         string
	OrganizationName string
	Email            string
	Password         string
}

type CreateOrganizationResponse struct {
}
