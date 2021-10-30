package gateway

type MessageRepository struct {
	DB *SqlHandler
}

func NewMessageRepository(db *SqlHandler) MessageRepository {
	return MessageRepository{DB: db}
}
