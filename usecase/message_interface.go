package usecase

import "github.com/ksrnnb/chat-app-server/entity"


type IMessageInteractor interface {
	SendMessage(*SendMessageRequest) (*SendMessageResponse, error)
}

type SendMessageRequest struct {
	RoomId int
	UserId int
	Text string
}

type SendMessageResponse struct {
	Message *entity.Message
}