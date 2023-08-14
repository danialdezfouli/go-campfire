package auth

type LoginInput struct {
	Email     string `json:"email" binding:"required"`
	Subdomain string `json:"subdomain" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
