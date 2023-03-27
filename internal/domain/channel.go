package domain

const ChannelTableName = "channels"

type Channel struct {
	Id             int
	OrganizationId int
	ProjectId      int
	Name           string
}

type ChannelRepository interface {
	CreateChannel(i Channel) *Channel
	UpdateChannel(i Channel) *Channel
	DeleteChannel(i Channel) *Channel
}
