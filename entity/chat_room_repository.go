package entity

type IChatRoomRepository interface {
	GetAllChatRooms(int) ([]*ChatRoom, error)
}
