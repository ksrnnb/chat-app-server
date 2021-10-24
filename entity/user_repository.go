package entity

type IUserRepository interface {
	GetUserByLoginId(string) (*User, error)
}
