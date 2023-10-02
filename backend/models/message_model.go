package models

import "time"

type Message struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"roomId"`
	Content   string    `json:"content"`
	SenderID  string    `json:"senderId"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewMessage(user_ID, room_ID, message, sender_ID string) *Message {
	return &Message{
		RoomID:    room_ID,
		Content:   message,
		SenderID:  sender_ID,
		CreatedAt: time.Now(),
	}
}
