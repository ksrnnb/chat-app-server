package usecase

import (
	"github.com/ksrnnb/chat-app-server/entity"
)

type MessageInteractor struct {
	MessageRepository entity.IMessageRepository
}

func NewMessageInteractor(r entity.IMessageRepository) *MessageInteractor {
	return &MessageInteractor{
		MessageRepository: r,
	}
}

func (interactor *MessageInteractor) GetMessages(req *GetMessagesRequest) (*GetMessagesResponse, error) {
	messages, err := interactor.MessageRepository.GetMessages(req.RoomId)

	if err != nil {
		return nil, err
	}

	res := &GetMessagesResponse{
		Messages: messages,
	}

	return res, nil
}