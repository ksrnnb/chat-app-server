package gateway

import (
	"github.com/ksrnnb/chat-app-server/entity"
	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *SqlHandler
}

func NewMessageRepository(db *SqlHandler) MessageRepository {
	return MessageRepository{DB: db}
}

// メッセージを取得する
func (r MessageRepository) GetMessages(roomId int) ([]*entity.Message, error) {
	var messages []*entity.Message
	err := r.DB.Debug().Preload("User", r.selectWithoutCredentials).Where("room_id = ?", roomId).Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

// 認証情報は取得しない
func (r MessageRepository) selectWithoutCredentials(db *gorm.DB) *gorm.DB {
	return db.Select("id, name, avatar")
}