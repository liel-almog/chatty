package repository

import (
	"chatty/backend/database"
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
)

type RoomRepository interface {
	FindAll() (pgx.Rows, error)
	FindOne(roomID string) pgx.Row
}

type RoomRepositoryImpl struct {
	db *database.PostgreSQLpgx
}

var (
	initRoomRepositoryOnce sync.Once
	roomRepository         *RoomRepositoryImpl
)

func (r *RoomRepositoryImpl) FindAll() (pgx.Rows, error) {
	rows, err := r.db.Pool.Query(context.Background(), "SELECT id, name, created_at from rooms")
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *RoomRepositoryImpl) FindOne(roomID string) pgx.Row {
	row := r.db.Pool.QueryRow(context.Background(), "SELECT id, name, created_at from rooms WHERE id = $1", roomID)

	return row
}

func InitRoomRepositoryImpl() *RoomRepositoryImpl {
	db := database.GetDB()

	return &RoomRepositoryImpl{
		db: db,
	}
}

func GetRoomRepository() RoomRepository {
	initRoomRepositoryOnce.Do(func() {
		roomRepository = InitRoomRepositoryImpl()
	})

	return roomRepository
}
