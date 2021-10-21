package usecase

import (
	"github.com/ksrnnb/chat-app-server/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserInteractor struct {
	UserRepository entity.IUserRepository
}

func NewUserInteractor(u entity.IUserRepository) UserInteractor {
	return UserInteractor{UserRepository: u}
}

func (u UserInteractor) GetUserByLoginId(req *LoginInput) (*LoginOutput, error) {
	user, err := u.UserRepository.GetUserByLoginId(req.LoginId)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, err
	}

	return &LoginOutput{User: user}, nil
}