package entity

type ChatRoom struct {
	Id       int
	Name  string
	Users []*User
	Messages []*Message
}
