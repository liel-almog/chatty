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

func NewMessage(user_id, room_id, message, sender_id string) *Message {
	return &Message{
		User_id:    user_id,
		Room_id:    room_id,
		Message:    message,
		Sender_id:  sender_id,
		Created_at: time.Now(),
	}
}
