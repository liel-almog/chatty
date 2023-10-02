package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUser(name string) *User {
	return &User{
		Name:      name,
		CreatedAt: time.Now(),
	}
}
