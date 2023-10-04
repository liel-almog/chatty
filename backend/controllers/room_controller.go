package controllers

import (
	"chatty/backend/service"
	"sync"

	"github.com/gin-gonic/gin"
)

type RoomController interface {
	GetAll(c *gin.Context)
}

type RoomControllerImpl struct {
	roomService service.RoomService
}

var (
	initRoomControllerOnce sync.Once
	roomController         *RoomControllerImpl
)

func (r *RoomControllerImpl) GetAll(c *gin.Context) {
	rooms, err := roomController.roomService.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, rooms)
}

func InitRoomControllerImpl() {
	roomService := service.GetRoomService()

	roomController = &RoomControllerImpl{
		roomService: roomService,
	}
}

func GetRoomController() RoomController {
	initRoomControllerOnce.Do(func() {
		InitRoomControllerImpl()
	})

	return roomController
}
