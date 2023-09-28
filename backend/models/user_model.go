package models

import "time"

type User struct {
	Id         string
	Name       string
	Created_at time.Time
	Message
}
