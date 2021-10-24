package usecase

import "github.com/ksrnnb/chat-app-server/entity"

type IUserInteractor interface {
	GetUserByLoginId(*LoginInput) (*LoginOutput, error)
}

type LoginInput struct {
	LoginId  string
	Password string
}

type LoginOutput struct {
	User *entity.User
}
