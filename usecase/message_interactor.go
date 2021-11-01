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

func (interactor MessageInteractor) SendMessage(req *SendMessageRequest) (*SendMessageResponse, error) {
	message, err := interactor.MessageRepository.SendMessage(req.RoomId, req.UserId, req.Text)

	if err != nil {
		return nil, err
	}

	res := &SendMessageResponse{
		Message: message,
	}

	return res, nil
}