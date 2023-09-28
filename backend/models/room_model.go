package models

import "time"

type Room struct {
	Id         string
	Name       string
	Created_at time.Time
}
