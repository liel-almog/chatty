package models

import "time"

type Room struct {
	Id         string
	Name       string
	Created_at time.Time
}

func NewRoom(name string) *Room {
	return &Room{
		Name:       name,
		Created_at: time.Now(),
	}
}
