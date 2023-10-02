package websocket

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
			for client := range h.Clients {
				// We don't want to send the message to the sender
				if client != broadcaster.Sender {
					select {
					case client.send <- broadcaster.Bytes:
					default:
						close(client.send)
						delete(h.Clients, client)
					}
				}
			}
		}
	}
}
