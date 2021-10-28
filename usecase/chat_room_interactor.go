package usecase

import (
	"github.com/ksrnnb/chat-app-server/entity"
)

type ChatRoomInteractor struct {
	ChatRoomRepository entity.IChatRoomRepository
}

func NewChatRoomInteractor(r entity.IChatRoomRepository) *ChatRoomInteractor {
	return &ChatRoomInteractor{
		ChatRoomRepository: r,
	}
}

func (interactor *ChatRoomInteractor) GetAllChatRooms(req *GetChatRoomsRequest) (*GetChatRoomsResponse, error) {
	rooms, err := interactor.ChatRoomRepository.GetAllChatRooms(req.UserId)

	if err != nil {
		return nil, err
	}

	res := &GetChatRoomsResponse{
		ChatRooms: rooms,
	}

	return res, nil
}