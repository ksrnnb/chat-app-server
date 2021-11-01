package request

type SendMessageRequest struct {
	Message string `json:"message"`
	RoomId int
	UserId int
}
