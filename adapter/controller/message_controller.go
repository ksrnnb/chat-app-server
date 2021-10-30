package controller

import (
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/response"
	"github.com/ksrnnb/chat-app-server/entity"
	"github.com/ksrnnb/chat-app-server/usecase"
)

type MessageController struct{}

func NewMessageController() MessageController {
	return MessageController{}
}

type Message struct {
	Id int
	RoomId int
	UserId int
	Text string
	User *User
}

type User struct {
	Id int
	Name string
	Avatar string
}

// メッセージを取得する
func (c MessageController) GetMessages(interactor usecase.IMessageInteractor, roomId int) *response.Response {
	req := &usecase.GetMessagesRequest{RoomId: roomId}
	res, err := interactor.GetMessages(req)

	if err != nil {
		return &response.Response{Code: http.StatusNotFound, Params: nil}
	}

	messages := c.removeCredentials(res.Messages)
	return &response.Response{Code: http.StatusOK, Params: messages}
}

// ログインIDとパスワードを除く
func (c MessageController) removeCredentials(messages []*entity.Message) []*Message {
	var entityMessages []*Message

	for _, message := range messages {
		newUser := &User{
			Id: message.User.Id,
			Name: message.User.Name,
			Avatar: message.User.Avatar,
		}

		newEntityMessage := &Message{
			Id: message.Id,
			RoomId: message.RoomId,
			UserId: message.UserId,
			Text: message.Text,
			User: newUser,
		}

		entityMessages = append(entityMessages, newEntityMessage)
	}

	return entityMessages
}

