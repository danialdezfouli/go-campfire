package postgresql

import (
	"campfire/internal/database"
	"campfire/internal/domain"
	"context"
)

type OrganizationRepositoryPostgres struct {
}

func NewOrganizationRepositoryPostgres() OrganizationRepositoryPostgres {
	return OrganizationRepositoryPostgres{}
}

func (o OrganizationRepositoryPostgres) CreateOrganization(ctx context.Context, input *domain.Organization) error {
	db := database.GetPostgres()
	sql := `INSERT INTO organizations (name, subdomain, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id`
	err := db.QueryRow(sql, input.Name, input.Subdomain).Scan(&input.Id)
	if err != nil {
		return err
	}

	return nil
}
