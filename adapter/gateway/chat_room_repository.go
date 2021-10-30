package gateway

import (
	"github.com/ksrnnb/chat-app-server/entity"
	"gorm.io/gorm"
)

type ChatRoomRepository struct {
	DB *SqlHandler
}

type ChatRoom struct {
	Id    int
	Name  string
	Users []*entity.User `gorm:"many2many:room_user;joinForeignKey:room_id"`
	Messages []*entity.Message `gorm:"foreignKey:room_id"`
}

type RoomUser struct {
	RoomId int
	UserId int
}

func (ru RoomUser) TableName() string {
	return "room_user"
}

func NewChatRoomRepository(db *SqlHandler) ChatRoomRepository {
	return ChatRoomRepository{DB: db}
}

// ユーザーに関連するルームを取得する
func (r ChatRoomRepository) GetAllChatRooms(userId int) ([]*entity.ChatRoom, error) {
	roomIds, err := r.getRoomIdsByUserId(userId)

	if err != nil {
		return nil, err
	}

	return r.getChatRoomsWithUsers(roomIds)
}

func (r ChatRoomRepository) GetChatRoom(roomId int) (*entity.ChatRoom, error) {
	var room *ChatRoom
	err := r.DB.
			Preload("Messages.User").
			Where("id = ?", roomId).
			Find(&room).
			Error

	if err != nil {
		return nil, err
	}

	entityRoom := toEntityChatRoom(room)

	return entityRoom, nil
}

// roomIdを取得する
func (r ChatRoomRepository) getRoomIdsByUserId(userId int) ([]int, error) {
	var roomUsers []*RoomUser

	err := r.DB.Select("room_id").Where("user_id = ?", userId).Find(&roomUsers).Error
	
	if err != nil {
		return nil, err
	}

	var roomIds []int
	for _, roomUser := range roomUsers {
		roomIds = append(roomIds, roomUser.RoomId)
	}

	return roomIds, nil
}

// roomIdからルームと関連するユーザーを取得
func (r ChatRoomRepository) getChatRoomsWithUsers(roomIds []int) ([]*entity.ChatRoom, error) {
	var rooms []*ChatRoom
	err := r.DB.Preload("Messages.User").
				Preload("Users").
				Where(roomIds).
				Find(&rooms).
				Error

	if err != nil {
		return nil, err
	}

	entityRooms := toEntityChatRooms(rooms)

	return entityRooms, nil
}

// 認証情報は取得しない
func (r ChatRoomRepository) selectWithoutCredentials(db *gorm.DB) *gorm.DB {
	return db.Select("id, name, avatar")
}

// gatewayのChatRoomからentityのChatRoomへ変換
func toEntityChatRooms(chatRooms []*ChatRoom) (entityChatRooms []*entity.ChatRoom) {
	for _, chatRoom := range chatRooms {
		newEntityChatRoom := toEntityChatRoom(chatRoom)

		entityChatRooms = append(entityChatRooms, newEntityChatRoom)
	}

	return entityChatRooms
}

// gatewayのChatRoomからentityのChatRoomへ変換
func toEntityChatRoom(chatRoom *ChatRoom) (*entity.ChatRoom) {
	entityChatRoom := &entity.ChatRoom{
		Id: chatRoom.Id,
		Name: chatRoom.Name,
		Users: chatRoom.Users,
		Messages: chatRoom.Messages,
	}

	return entityChatRoom
}