package gateway

import "github.com/ksrnnb/chat-app-server/entity"

type MessageRepository struct {
	DB *SqlHandler
}

func NewMessageRepository(db *SqlHandler) MessageRepository {
	return MessageRepository{DB: db}
}

func (r MessageRepository) SendMessage(roomId int, userId int, text string) (*entity.Message, error) {
	entityMessage := &entity.Message{
		RoomId: roomId,
		UserId: userId,
		Text: text,
	}

	err := r.DB.Create(entityMessage).Error

	if err != nil {
		return nil, err
	}

	return entityMessage, nil
}