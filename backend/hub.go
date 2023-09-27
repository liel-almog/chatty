package main

// This is a struct that holds the bytes to be broadcasted and the sender
type Broadcaster struct {
	bytes  []byte
	sender *Client
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *Broadcaster

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Broadcaster),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case broadcaster := <-h.broadcast:
			for client := range h.clients {
				// We don't want to send the message to the sender
				if client != broadcaster.sender {
					select {
					case client.send <- broadcaster.bytes:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
