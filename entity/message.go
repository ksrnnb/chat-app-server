package entity

type Message struct {
	Id       int
	RoomId int
	UserId int
	Text  string
	User *User
}
