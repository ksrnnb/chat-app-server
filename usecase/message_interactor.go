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
