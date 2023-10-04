package repository

import (
	"chatty/backend/database"
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
)

type RoomRepository interface {
	FindAll() (pgx.Rows, error)
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
