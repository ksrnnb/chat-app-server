package repository

import (
	"github.com/ksrnnb/chat-app-server/entity"
)

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (u UserRepository) GetUserByLoginId(loginId string) (*entity.User, error) {
	// hoge := &database.DBClient{}

	return &entity.User{}, nil
}