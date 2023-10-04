package service

import (
	"chatty/backend/models"
	"chatty/backend/repository"
	"sync"
)

type RoomService interface {
	GetAll() (*[]models.Room, error)
}

type RoomServiceImpl struct {
	roomRepository repository.RoomRepository
}

var (
	initMessageServiceOnce sync.Once
	roomService            *RoomServiceImpl
)

func (r *RoomServiceImpl) GetAll() (*[]models.Room, error) {
	rows, err := r.roomRepository.FindAll()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room

	for rows.Next() {
		var room models.Room
		err = rows.Scan(&room.ID, &room.Name, &room.CreatedAt)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &rooms, nil
}

func InitRoomServiceImpl() *RoomServiceImpl {
	roomRepository := repository.GetRoomRepository()

	return &RoomServiceImpl{
		roomRepository: roomRepository,
	}
}

func GetRoomService() RoomService {
	initMessageServiceOnce.Do(func() {
		roomService = InitRoomServiceImpl()
	})

	return roomService
}
