package models

import "time"

type Message struct {
	Id         string
	User_id    string
	Room_id    string
	Message    string
	Sender_id  string
	Created_at time.Time
}
