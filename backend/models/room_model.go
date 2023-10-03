package models

import "time"

type Room struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewRoom(name string) *Room {
	return &Room{
		Name:      name,
		CreatedAt: time.Now(),
	}
}
