package controllers

import (
	"chatty/backend/service"

	"github.com/gin-gonic/gin"
)

type MessageController interface {
	// GetAll a new message
	GetAll(c *gin.Context)
}

type MessageControllerImpl struct {
	messageService service.MessageService
}

var messageController *MessageControllerImpl

func (m *MessageControllerImpl) GetAll(c *gin.Context) {
	messages, err := messageController.messageService.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": messages,
	})
}

func InitMessageControllerImpl() {
	messageService := service.GetMessageService()

	messageController = &MessageControllerImpl{
		messageService: messageService,
	}
}

func GetMessageController() MessageController {
	// Check if messageController is nil
	if messageController == nil {
		InitMessageControllerImpl()
	}

	return messageController
}
