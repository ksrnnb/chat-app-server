package controller

import (
	"net/http"

	"github.com/ksrnnb/chat-app-server/adapter/response"
	"github.com/ksrnnb/chat-app-server/usecase"
)

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

func (c UserController) GetUser(session ISession, interactor usecase.IUserInteractor) *response.Response {
	id, ok := session.Get("userId").(int)

	if !ok {
		return &response.Response{
			Code: http.StatusNotFound,
		}
	}

	req := &usecase.FindInput{Id: id}

	res, err := interactor.Find(req)

	if err != nil {
		return &response.Response{
			Code: http.StatusNotFound,
		}
	}

	return &response.Response{
		Code: http.StatusOK,
		Params: map[string]interface{}{
			"id": res.User.Id,
		},
	}
}