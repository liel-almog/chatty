package repository

import (
	"chatty/backend/database"
	"chatty/backend/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type MessageRepository interface {
	FindAll() (pgx.Rows, error)
	Save(message *models.CreateMessageDTO) error
}

type MessageRepositoryImpl struct {
	db *database.PostgreSQLpgx
}

var messageRepository *MessageRepositoryImpl

func (m *MessageRepositoryImpl) FindAll() (pgx.Rows, error) {
	rows, err := m.db.Pool.Query(context.Background(), "SELECT * from messages")
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (m *MessageRepositoryImpl) Save(message *models.CreateMessageDTO) error {
	_, err :=
		m.db.Pool.Exec(context.Background(),
			"INSERT INTO messages (room_id, content) VALUES ($1, $2)",
			message.RoomID, message.Content)

	if err != nil {
		return err
	}

	return nil
}

func InitMessageRepositoryImpl() {
	db := database.GetDB()

	messageRepository = &MessageRepositoryImpl{
		db: db,
	}
}

func GetMessageRepository() MessageRepository {
	// Check if messageRepository is nil
	if messageRepository == nil {
		InitMessageRepositoryImpl()
	}

	return messageRepository
}
