package repository

import (
	"campfire/internal/database"
	"campfire/internal/domain"
	"context"
	"fmt"
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

func (r UserRepositoryPostgres) GetUser(ctx context.Context, userId int) (*domain.User, error) {
	db := database.GetPostgres()
	sql := `SELECT name, email FROM users WHERE id=$1`
	var user domain.User
	err := db.QueryRow(sql, userId).Scan(&user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	user.Id = fmt.Sprintf("%d", userId)
	return &user, nil
}
