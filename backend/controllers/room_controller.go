package controllers

import (
	"chatty/backend/service"
	"sync"

	"github.com/gin-gonic/gin"
)

type RoomController interface {
	GetAll(c *gin.Context)
	GetAllMessagesByRoom(c *gin.Context)
}

type RoomControllerImpl struct {
	roomService    service.RoomService
	messageService service.MessageService
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

func (r *RoomControllerImpl) GetAllMessagesByRoom(c *gin.Context) {
	messages, err := r.messageService.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, messages)
}

func InitRoomControllerImpl() *RoomControllerImpl {
	roomService := service.GetRoomService()
	messageService := service.GetMessageService()

	return &RoomControllerImpl{
		roomService:    roomService,
		messageService: messageService,
	}
}

func GetRoomController() RoomController {
	initRoomControllerOnce.Do(func() {
		roomController = InitRoomControllerImpl()
	})

	return roomController
}
