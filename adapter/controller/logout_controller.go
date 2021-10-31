package controller

import (
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/response"
)

type LogoutController struct{}

func NewLogoutController() LogoutController {
	return LogoutController{}
}

func (c LogoutController) Logout(session ISession) *response.Response {
	err := session.Disable()

	if err != nil {
		return &response.Response{
			Code: http.StatusInternalServerError,
			Params: nil,
		}
	}

	return &response.Response{
		Code: http.StatusOK,
		Params: map[string]interface{}{
			"message": "ok",
		},
	}
}