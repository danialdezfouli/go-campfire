package domain

const TextMessageTableName = "messages"

type TextMessage struct {
	Id        int
	ChannelId int
	UserId    int
	Name      string
}

type TextReaction struct {
	Id            int
	UserId        int
	TextMessageId int
	Reaction      string
}

type TextMessageRepository interface {
	GetMessage(id int) *TextMessage
	ReactMessage(id int, react string) *TextMessage
	DeleteMessage(id int) error
}
