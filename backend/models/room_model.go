package models

import "time"

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewRoom(name string) *Room {
	return &Room{
		Name:      name,
		CreatedAt: time.Now(),
	}
}
