package service

import (
	"chatty/backend/models"
	"chatty/backend/repository"
)

type MessageService interface {
	// GetAll a new message
	GetAll() ([]models.Message, error)
}

type MessageServiceImpl struct {
	messageRepository repository.MessageRepository
}

var messageService *MessageServiceImpl

func (m *MessageServiceImpl) GetAll() ([]models.Message, error) {
	rows, err := m.messageRepository.FindAll()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		var message models.Message
		err = rows.Scan(&message.ID, &message.RoomID, &message.Content, &message.SenderID, &message.CreatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func InitMessageServiceImpl() {
	messageRepository := repository.GetMessageRepository()

	messageService = &MessageServiceImpl{
		messageRepository: messageRepository,
	}
}

func GetMessageService() MessageService {
	// Check if messageService is nil
	if messageService == nil {
		InitMessageServiceImpl()
	}

	return messageService
}
