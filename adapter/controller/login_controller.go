package controller

import (
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/request"
	"github.com/ksrnnb/chat-app-server/adapter/response"
	"github.com/ksrnnb/chat-app-server/usecase"
)

type LoginController struct{}

func NewLoginController() LoginController {
	return LoginController{}
}

func (c LoginController) Login(req request.LoginRequest, interactor usecase.IUserInteractor) *response.Response {
	in := &usecase.LoginInput{
		LoginId: req.LoginId,
		Password: req.Password,
	}

	res, err := interactor.GetUserByLoginId(in)

	if err != nil {
		return &response.Response{
			Code: http.StatusUnauthorized,
			Params: nil,
		}
	}

	return &response.Response{
		Code: http.StatusOK,
		Params: map[string]interface{}{
			"id": res.User.Id,
		},
	}
}