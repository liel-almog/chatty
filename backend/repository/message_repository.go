package repository

import (
	"chatty/backend/database"
	"chatty/backend/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type MessageRepository interface {
	FindAll() (pgx.Rows, error)
	Save(message *models.CreateMessageDTO) (*models.Message, error)
}

type MessageRepositoryImpl struct {
	db *database.PostgreSQLpgx
}

var messageRepository *MessageRepositoryImpl

func (m *MessageRepositoryImpl) FindAll() (pgx.Rows, error) {
	rows, err := m.db.Pool.Query(context.Background(), "SELECT id, room_id, content, created_at from messages")
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (m *MessageRepositoryImpl) Save(message *models.CreateMessageDTO) (*models.Message, error) {
	var newMessage models.Message

	err :=
		m.db.Pool.QueryRow(context.Background(),
			"INSERT INTO messages (room_id, content) VALUES ($1, $2) RETURNING id, room_id, content, created_at",
			message.RoomID, message.Content,
		).Scan(&newMessage.ID, &newMessage.RoomID, &newMessage.Content, &newMessage.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &newMessage, nil
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
