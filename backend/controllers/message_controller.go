package controllers

import (
	"chatty/backend/service"
	"sync"

	"github.com/gin-gonic/gin"
)

type MessageController interface {
	GetAll(c *gin.Context)
}

type MessageControllerImpl struct {
	messageService service.MessageService
}

var (
	initMessageConrollerOnce sync.Once
	messageController        *MessageControllerImpl
)

func (m *MessageControllerImpl) GetAll(c *gin.Context) {
	messages, err := messageController.messageService.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, messages)
}

func InitMessageControllerImpl() *MessageControllerImpl {
	messageService := service.GetMessageService()

	return &MessageControllerImpl{
		messageService: messageService,
	}
}

func GetMessageController() MessageController {
	initMessageConrollerOnce.Do(func() {
		messageController = InitMessageControllerImpl()
	})

	return messageController
}
