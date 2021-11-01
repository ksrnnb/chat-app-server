package controller

import (
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/request"
	"github.com/ksrnnb/chat-app-server/adapter/response"
	"github.com/ksrnnb/chat-app-server/usecase"
)

type MessageController struct{}

func NewMessageController() MessageController {
	return MessageController{}
}

// チャットメッセージの登録
func (c MessageController) SendMessage(req request.SendMessageRequest, interactor usecase.IMessageInteractor) *response.Response {
	usecaseReq := &usecase.SendMessageRequest{
		RoomId: req.RoomId,
		UserId: req.UserId,
		Text: req.Message,
	}

	res, err := interactor.SendMessage(usecaseReq)

	if err != nil {
		return &response.Response{
			Code: http.StatusInternalServerError,
			Params: nil,
		}
	}

	newRes := map[string]interface{}{
		"id": res.Message.Id,
	}

	return &response.Response{
		Code: http.StatusCreated,
		Params: newRes,
	}
}