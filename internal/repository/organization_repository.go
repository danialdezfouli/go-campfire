package repository

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

func (o OrganizationRepositoryPostgres) FindById(ctx context.Context, id int) (*domain.Organization, error) {
	db := database.GetPostgres()

	rows, err := db.Query("SELECT id, name FROM organizations WHERE id = $1 limit 1", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		m := &domain.Organization{}
		err = rows.Scan(&m.Id, &m.Name)
		if err != nil {
			return nil, err
		}
		return m, nil
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nil, nil
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
