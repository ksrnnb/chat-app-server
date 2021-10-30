package controller

import (
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/response"
	"github.com/ksrnnb/chat-app-server/entity"
	"github.com/ksrnnb/chat-app-server/usecase"
)

type ChatRoomController struct {}

type ChatRoom struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Users []*User `json:"users"`
	Messages []*Message `json:"messages"`
}

type Message struct {
	Id int `json:"id"`
	RoomId int `json:"roomId"`
	UserId int `json:"userId"`
	Text string `json:"text"`
	User *User `json:"user"`
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}

func NewChatRoomController() ChatRoomController {
	return ChatRoomController{}
}

func (c ChatRoomController) GetChatRooms(interactor usecase.IChatRoomInteractor, userId int) *response.Response {
	req := &usecase.GetChatRoomsRequest{UserId: userId}
	res, err := interactor.GetAllChatRooms(req)

	if err != nil {
		return &response.Response{Code: http.StatusNotFound, Params: nil}
	}

	params := c.toSuitableChatRoomsFormat(res.ChatRooms)

	return &response.Response{Code: http.StatusOK, Params: params}
}

func (c ChatRoomController) GetRoom(interactor usecase.IChatRoomInteractor, roomId int) *response.Response {
	req := &usecase.GetChatRoomRequest{RoomId: roomId}
	res, err := interactor.GetChatRoom(req)

	if err != nil {
		return &response.Response{Code: http.StatusNotFound, Params: nil}
	}

	room := c.toSuitableChatRoomFormat(res.Room)
	return &response.Response{Code: http.StatusOK, Params: room}
}

// ルームの構造体から必要なフィールドだけ抽出
// エンティティ層にJSONの記述はしたくない
// 構造体のままだと以下のようなレスポンスとなる
// {
// "ChatRooms": [
//     {
//         "Id": 1,
//         "Name": "chat room 1",
//         "Users": [
//             {
//                 ...
//             },
func (c ChatRoomController) toSuitableChatRoomsFormat(chatRooms []*entity.ChatRoom) []*ChatRoom {
	var res []*ChatRoom

	for _, room := range chatRooms {
		formatRoom := c.toSuitableChatRoomFormat(room)
		res = append(res, formatRoom)
	}

	return res
}

// チャットルームの変換
func (c ChatRoomController) toSuitableChatRoomFormat(room *entity.ChatRoom) *ChatRoom {
	var newMessages []*Message

	for _, message := range room.Messages {
		newUser := c.toSuitableUserFormat(message.User)

		newMessage := &Message{
			Id: message.Id,
			RoomId: message.RoomId,
			UserId: message.UserId,
			Text: message.Text,
			User: newUser,
		}

		newMessages = append(newMessages, newMessage)
	}

	users := c.toSuitableUsersFormat(room.Users)
	newRoom := &ChatRoom{
		Id: room.Id,
		Name: room.Name,
		Messages: newMessages,
		Users: users,
	}

	return newRoom
}

// ユーザーの構造体から必要なフィールドだけ抽出
func (c ChatRoomController) toSuitableUsersFormat(users []*entity.User) []*User {
	var res []*User

	for _, user := range users {
		formatUser := c.toSuitableUserFormat(user)
		res = append(res, formatUser)
	}

	return res
}

func (c ChatRoomController) toSuitableUserFormat(entityUser *entity.User) *User {
	if entityUser == nil {
		return nil
	}

	newUser := &User{
		Id: entityUser.Id,
		Name: entityUser.Name,
		Avatar: entityUser.Avatar,
	}

	return newUser
}
