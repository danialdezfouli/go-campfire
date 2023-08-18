package domain

import "context"

const OrganizationTableName = "organizations"

type Organization struct {
	Id        int    `bson:"id,omitempty"`
	Name      string `bson:"name,omitempty"`
	Subdomain string `bson:"subdomain,omitempty"`
}

type OrganizationRepository interface {
	FindById(ctx context.Context, id int) (*Organization, error)
	CreateOrganization(ctx context.Context, i *Organization) error
	DeleteOrganization(ctx context.Context, id int) error
}
