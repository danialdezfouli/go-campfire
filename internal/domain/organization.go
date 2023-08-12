package domain

import "context"

const OrganizationTableName = "organizations"

type Organization struct {
	Id        interface{} `bson:"id,omitempty"`
	Name      string      `bson:"name,omitempty"`
	Subdomain string      `bson:"subdomain,omitempty"`
}

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, i *Organization) error
	GetOrganization(ctx context.Context, organizationId int) (*Organization, error)
}
