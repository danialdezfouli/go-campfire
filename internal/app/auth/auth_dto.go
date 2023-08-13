package auth

type LoginInput struct {
	Email     string `form:"email" validate:"required"`
	Subdomain string `form:"subdomain" validate:"required"`
	Password  string `form:"password" validate:"required"`
}
