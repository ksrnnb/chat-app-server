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

	u.DB.Where("login_id = ?", loginId).First(&user)
	
	err := u.DB.Error()

	if err != nil {
		return nil, err
	}

	return &user, nil
}