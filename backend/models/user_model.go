package models

import "time"

type User struct {
	Id         string
	Name       string
	Created_at time.Time
}

func NewUser(name string) *User {
	return &User{
		Name:       name,
		Created_at: time.Now(),
	}
}
