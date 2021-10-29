package controller

import (
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/response"
	"github.com/ksrnnb/chat-app-server/entity"
	"github.com/ksrnnb/chat-app-server/usecase"
)

type ChatRoomController struct {}

func NewChatRoomController() ChatRoomController {
	return ChatRoomController{}
}

func (c ChatRoomController) GetChatRooms(interactor usecase.IChatRoomInteractor, userId int) *response.Response {
	req := &usecase.GetChatRoomsRequest{UserId: userId}
	res, err := interactor.GetAllChatRooms(req)

	if err != nil {
		return &response.Response{Code: http.StatusNotFound, Params: nil}
	}

	params := c.toSuitableFormat(res.ChatRooms)

	return &response.Response{Code: http.StatusOK, Params: params}
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
func (c ChatRoomController) toSuitableFormat(chatRooms []*entity.ChatRoom) []map[string]interface{} {
	var res []map[string]interface{}

	for _, room := range chatRooms {
		formatRoom := make(map[string]interface{})
		users := c.toSuitableUsersFormat(room.Users)
		formatRoom["id"] = room.Id
		formatRoom["name"] = room.Name
		formatRoom["users"] = users

		res = append(res, formatRoom)
	}

	return res
}

// ユーザーの構造体から必要なフィールドだけ抽出
func (c ChatRoomController) toSuitableUsersFormat(users []*entity.User) []map[string]interface{} {
	var res []map[string]interface{}

	for _, user := range users {
		formatUser := make(map[string]interface{})
		formatUser["id"] = user.Id
		formatUser["name"] = user.Name
		formatUser["avatar"] = user.Avatar
		res = append(res, formatUser)
	}

	return res
}