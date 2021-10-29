package gateway

import (
	"github.com/ksrnnb/chat-app-server/entity"
)

type ChatRoomRepository struct {
	DB *SqlHandler
}

type ChatRoom struct {
	Id    int
	Name  string
	Users []*entity.User `gorm:"many2many:room_user;joinForeignKey:room_id"`
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
	err := r.DB.Preload("Users").Where(roomIds).Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	entityRooms := toEntityChatRooms(rooms)

	return entityRooms, nil
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