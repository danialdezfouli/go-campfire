package membership

type AddMemberInput struct {
	OrganizationId int    `validate:"required"`
	UserName       string `validate:"required"`
	Email          string `validate:"required"`
	Password       string `validate:"required"`
}
