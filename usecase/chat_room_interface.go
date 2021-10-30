package usecase

import "github.com/ksrnnb/chat-app-server/entity"


type IChatRoomInteractor interface {
	GetAllChatRooms(*GetChatRoomsRequest) (*GetChatRoomsResponse, error)
	GetChatRoom(*GetChatRoomRequest) (*GetChatRoomResponse, error)
}

type GetChatRoomsRequest struct {
	UserId int
}

type GetChatRoomsResponse struct {
	ChatRooms []*entity.ChatRoom
}

type GetChatRoomRequest struct {
	RoomId int
}

type GetChatRoomResponse struct {
	Room *entity.ChatRoom
}