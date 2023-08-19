package domain

import "context"

const OrganizationTableName = "organizations"

type OrganizationId uint
type Organization struct {
	Id        OrganizationId `bson:"id,omitempty"`
	Name      string         `bson:"name,omitempty"`
	Subdomain string         `bson:"subdomain,omitempty"`
}

type OrganizationRepository interface {
	FindById(ctx context.Context, id OrganizationId) (*Organization, error)
	CreateOrganization(ctx context.Context, i *Organization) error
	DeleteOrganization(ctx context.Context, id OrganizationId) error
}
