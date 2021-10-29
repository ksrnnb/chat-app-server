package entity

type IMessageRepository interface {
	GetMessages(int) ([]*Message, error)
}
