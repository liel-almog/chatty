package models

import (
	"chatty/backend/configs"
	"encoding/json"
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"roomId" validate:"required"`
	Content   string    `json:"content" validate:"required,max=500"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username" validate:"required,max=50"`
}

type CreateMessageDTO struct {
	RoomID   int    `json:"roomId" validate:"required"`
	Content  string `json:"content" validate:"required,max=500"`
	Username string `json:"username" validate:"required,max=50"`
}

func (m *CreateMessageDTO) IsValid() (bool, error) {
	validate := configs.GetValidator()

	err := validate.Struct(m)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *CreateMessageDTO) UnmarshalCreateMessage(data []byte) (*CreateMessageDTO, error) {
	message := &CreateMessageDTO{}
	err := json.Unmarshal(data, message)

	return message, err
}
