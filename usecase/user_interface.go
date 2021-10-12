package usecase

import "github.com/ksrnnb/chat-app-server/entity"

type IUserInteractor interface {
	GetUserByLoginId(*LoginRequest) (*LoginResponse, error)
}

type LoginRequest struct {
	LoginId string
	Password string
}

type LoginResponse struct {
	User *entity.User
}