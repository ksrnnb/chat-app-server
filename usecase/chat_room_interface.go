package usecase

import "github.com/ksrnnb/chat-app-server/entity"


type IChatRoomInteractor interface {
	GetAllChatRooms(*GetChatRoomsRequest) (*GetChatRoomsResponse, error)
}

type GetChatRoomsRequest struct {
	UserId int
}

type GetChatRoomsResponse struct {
	ChatRooms []*entity.ChatRoom
}