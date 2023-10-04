package websocket

import (
	"chatty/backend/models"
	"chatty/backend/repository"
	"encoding/json"
	"fmt"
)

// This is a struct that holds the bytes to be broadcasted and the sender
type Broadcaster struct {
	Bytes  []byte
	Sender *Client
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan *Broadcaster

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan *Broadcaster),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) processMessage(broadcaster *Broadcaster) {
	createMessage := &models.CreateMessageDTO{}
	createMessage, err := createMessage.UnmarshalCreateMessage(broadcaster.Bytes)

	if err != nil {
		fmt.Println("Error unmarshaling message:", err)
		return
	}

	// Validate the message
	if ok, err := createMessage.IsValid(); !ok {
		fmt.Println("Invalid message:", err)
		return
	}

	// Save the message
	row := repository.GetMessageRepository().Save(createMessage)

	var message models.Message
	err = row.Scan(&message.ID, &message.RoomID, &message.Content, &message.CreatedAt)

	if err != nil {
		fmt.Println("Error saving message:", err)
		return
	}

	// Broadcast the message
	broadcaster.Bytes, err = json.Marshal(message)

	if err != nil {
		fmt.Println("Error marshaling message:", err)
		return
	}

	h.broadcastMessage(broadcaster)
}

// Broadcast the message to all clients except the sender
func (h *Hub) broadcastMessage(broadcaster *Broadcaster) {
	for client := range h.Clients {
		select {
		case client.send <- broadcaster.Bytes:
		default:
			close(client.send)
			delete(h.Clients, client)
		}

	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}

		case broadcaster := <-h.Broadcast:
			h.processMessage(broadcaster)
		}
	}
}
