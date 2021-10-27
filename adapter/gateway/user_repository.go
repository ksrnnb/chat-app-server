package gateway

import "github.com/ksrnnb/chat-app-server/entity"

type UserRepository struct {
	DB *SqlHandler
}

func NewUserRepository(db *SqlHandler) UserRepository {
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

func (u UserRepository) Find(id int) (*entity.User, error) {
	var user entity.User

	err := u.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}