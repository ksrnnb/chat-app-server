package entity

type IMessageRepository interface {
	SendMessage(roomId int, userId int, message string) (*Message, error)
}
