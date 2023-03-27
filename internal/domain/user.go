package domain

import "context"

const UserTableName = "user"

type User struct {
	Id             string
	OrganizationId interface{}
	IsSuperAdmin   bool
	Name           string
	Email          string
	Password       string
	CreatedAt      string
	UpdatedAt      string
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, userId int) (*User, error)
}
