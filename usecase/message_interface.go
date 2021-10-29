package usecase

import "github.com/ksrnnb/chat-app-server/entity"


type IMessageInteractor interface {
	GetMessages(*GetMessagesRequest) (*GetMessagesResponse, error)
}

type GetMessagesRequest struct {
	RoomId int
}

type GetMessagesResponse struct {
	Messages []*entity.Message
}