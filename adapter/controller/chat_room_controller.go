package controller

import (
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/response"
	"github.com/ksrnnb/chat-app-server/usecase"
)

type ChatRoomController struct {}

func NewChatRoomController() ChatRoomController {
	return ChatRoomController{}
}

func (c ChatRoomController) GetChatRooms(interactor usecase.IChatRoomInteractor, userId int) *response.Response {
	req := &usecase.GetChatRoomsRequest{UserId: userId}
	chatRooms, err := interactor.GetAllChatRooms(req)

	if err != nil {
		return &response.Response{Code: http.StatusNotFound, Params: nil}
	}

	return &response.Response{Code: http.StatusOK, Params: chatRooms}
}