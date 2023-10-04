package service

import (
	"chatty/backend/models"
	"chatty/backend/repository"
	"sync"
)

type MessageService interface {
	// GetAll a new message
	GetAllByRoom(roomID string) (*[]models.Message, error)
	Create(message *models.CreateMessageDTO) (*models.Message, error)
}

type MessageServiceImpl struct {
	messageRepository repository.MessageRepository
}

var (
	initRoomServiceOnce sync.Once
	messageService      *MessageServiceImpl
)

func (m *MessageServiceImpl) GetAllByRoom(roomID string) (*[]models.Message, error) {
	rows, err := m.messageRepository.FindAllByRoom(roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		var message models.Message
		err = rows.Scan(&message.ID, &message.RoomID, &message.Content, &message.CreatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &messages, nil
}

func (m *MessageServiceImpl) Create(message *models.CreateMessageDTO) (*models.Message, error) {
	var newMessage models.Message

	row := m.messageRepository.Save(message)
	err := row.Scan(&newMessage.ID, &newMessage.RoomID, &newMessage.Content, &newMessage.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &newMessage, nil
}

func InitMessageServiceImpl() *MessageServiceImpl {
	messageRepository := repository.GetMessageRepository()

	return &MessageServiceImpl{
		messageRepository: messageRepository,
	}
}

func GetMessageService() MessageService {
	initRoomServiceOnce.Do(func() {
		messageService = InitMessageServiceImpl()
	})

	return messageService
}
