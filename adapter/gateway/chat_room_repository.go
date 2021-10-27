package gateway

import (
	"github.com/ksrnnb/chat-app-server/entity"
)

type ChatRoomRepository struct {
	DB *SqlHandler
}

type ChatRoom struct {
	entity.ChatRoom
	Users []*entity.User `gorm:"many2many:chat_user"`
}

func NewChatRoomRepository(db *SqlHandler) ChatRoomRepository {
	return ChatRoomRepository{DB: db}
}

func (r ChatRoomRepository) GetAllChatRooms(userId int) ([]*entity.ChatRoom, error) {
	var chatRooms []*ChatRoom
	err := r.DB.Preload("Users").Find(&chatRooms, "users.id = ?", userId).Error

	if err != nil {
		return nil, err
	}

	entityChatRooms := toEntityChatRooms(chatRooms)
	return entityChatRooms, nil
}

// gatewayのChatRoomからentityのChatRoomへ変換
func toEntityChatRooms(chatRooms []*ChatRoom) (entityChatRooms []*entity.ChatRoom) {
	for _, chatRoom := range chatRooms {
		newEntityChatRoom := &entity.ChatRoom{
			Id: chatRoom.Id,
			Name: chatRoom.Name,
			Users: chatRoom.Users,
		}

		entityChatRooms = append(entityChatRooms, newEntityChatRoom)
	}

	return entityChatRooms
}