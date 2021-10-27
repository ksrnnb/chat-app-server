package usecase

import "github.com/ksrnnb/chat-app-server/entity"

type ChatRoomInteractor struct {
	ChatRoomRepository entity.IChatRoomRepository
}

func NewChatRoomInteractor(r entity.IChatRoomRepository) *ChatRoomInteractor {
	return &ChatRoomInteractor{
		ChatRoomRepository: r,
	}
}

func (interactor *ChatRoomInteractor) GetAllChatRooms(*GetChatRoomsRequest) (*GetChatRoomsResponse, error) {
	return &GetChatRoomsResponse{}, nil
}