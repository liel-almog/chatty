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
}

type CreateMessage struct {
	RoomID  int    `json:"roomId" validate:"required"`
	Content string `json:"content" validate:"required,max=500"`
}

func (m *CreateMessage) IsValid() (bool, error) {
	validate := configs.GetValidator()

	err := validate.Struct(m)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *CreateMessage) UnmarshalCreateMessage(data []byte) (*CreateMessage, error) {
	message := &CreateMessage{}
	err := json.Unmarshal(data, message)

	return message, err
}
