package domain

import "context"

const UserTableName = "user"

type UserId = uint
type User struct {
	Id             UserId         `json:"id"`
	OrganizationId OrganizationId `json:"organization_id"`
	IsSuperAdmin   bool           `json:"is_super_admin"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Password       string         `json:"-"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserById(ctx context.Context, userId UserId) (*User, error)
	GetUserByEmail(ctx context.Context, email string, subdomain string) (*User, error)
}
