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
	// extract params
	roomID := c.Param("id")

	_, err := r.roomService.GetOne(roomID)

	if err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(404, gin.H{
				"error": "Room not found",
			})

			return
		}

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	messages, err := r.messageService.GetAllByRoom(roomID)

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
