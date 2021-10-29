package usecase

import "github.com/ksrnnb/chat-app-server/entity"

type IUserInteractor interface {
	GetUserByLoginId(*LoginInput) (*LoginOutput, error)
	Find(req *FindInput) (*FindOutput, error) 
}

type LoginInput struct {
	LoginId  string
	Password string
}

type LoginOutput struct {
	User *entity.User
}

type FindInput struct {
	Id int
}

type FindOutput struct {
	User *entity.User
}