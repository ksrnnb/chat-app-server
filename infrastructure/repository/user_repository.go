package repository

import (
	"github.com/ksrnnb/chat-app-server/entity"
	"github.com/ksrnnb/chat-app-server/infrastructure/database"
)

type UserRepository struct {
	DB *database.SqlHandler
}

func NewUserRepository(db *database.SqlHandler) UserRepository {
	return UserRepository{DB: db}
}

func (u UserRepository) GetUserByLoginId(loginId string) (*entity.User, error) {
	var user entity.User

	err := u.DB.Where("login_id = ?", loginId).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}