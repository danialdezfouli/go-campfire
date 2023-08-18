package auth

type LoginInput struct {
	Email     string `json:"email" validate:"required"`
	Subdomain string `json:"subdomain" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
