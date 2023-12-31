package repository

import (
	"chatty/backend/database"
	"chatty/backend/models"
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
)

type MessageRepository interface {
	FindAllByRoom(roomID string) (pgx.Rows, error)
	Save(message *models.CreateMessageDTO) pgx.Row
}

type MessageRepositoryImpl struct {
	db *database.PostgreSQLpgx
}

var (
	initMessageRepositoryOnce sync.Once
	messageRepository         *MessageRepositoryImpl
)

func (m *MessageRepositoryImpl) FindAllByRoom(roomID string) (pgx.Rows, error) {
	rows, err := m.db.Pool.Query(context.Background(),
		"SELECT id, room_id, content, created_at, username from messages WHERE room_id = $1 ORDER BY created_at ASC",
		roomID)

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (m *MessageRepositoryImpl) Save(message *models.CreateMessageDTO) pgx.Row {
	row :=
		m.db.Pool.QueryRow(context.Background(),
			"INSERT INTO messages (room_id, content, username) VALUES ($1, $2, $3) RETURNING id, room_id, content, created_at, username",
			message.RoomID, message.Content, message.Username,
		)

	return row
}

func InitMessageRepositoryImpl() *MessageRepositoryImpl {
	db := database.GetDB()

	return &MessageRepositoryImpl{
		db: db,
	}
}

func GetMessageRepository() MessageRepository {
	initMessageRepositoryOnce.Do(func() {
		messageRepository = InitMessageRepositoryImpl()
	})

	return messageRepository
}
