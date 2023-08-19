package repository

import (
	"campfire/internal/database"
	"campfire/internal/domain"
	"context"
)

type UserRepositoryPostgres struct {
}

func NewUserRepositoryPostgres() UserRepositoryPostgres {
	return UserRepositoryPostgres{}
}

func (r UserRepositoryPostgres) CreateUser(ctx context.Context, user *domain.User) error {
	db := database.GetPostgres()
	sql := `INSERT INTO users (name, email, password, organization_id, is_super_admin, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`
	err := db.QueryRow(sql, user.Name, user.Email, user.Password, user.OrganizationId, user.IsSuperAdmin).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepositoryPostgres) GetUserById(ctx context.Context, userId domain.UserId) (*domain.User, error) {
	db := database.GetPostgres()
	sql := `SELECT id, name, email, organization_id, is_super_admin, created_at, updated_at FROM users WHERE id=$1`
	var user domain.User
	err := db.QueryRow(sql, userId).
		Scan(&user.Id, &user.Name, &user.Email, &user.OrganizationId, &user.IsSuperAdmin, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r UserRepositoryPostgres) GetUserByEmail(ctx context.Context, email string, subdomain string) (*domain.User, error) {
	db := database.GetPostgres()
	sql := `SELECT users.id, users.name, users.organization_id, users.email, password FROM users
	inner join organizations on users.organization_id = organizations.id
	WHERE email=$1 and organizations.subdomain=$2`

	var user domain.User
	err := db.QueryRow(sql, email, subdomain).
		Scan(&user.Id, &user.Name, &user.OrganizationId, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
