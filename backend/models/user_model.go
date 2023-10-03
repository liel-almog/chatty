package models

import "time"

type User struct {
	ID        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(name string) *User {
	return &User{
		Name:      name,
		CreatedAt: time.Now(),
	}
}
