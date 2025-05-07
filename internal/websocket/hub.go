package websocket

type Hub struct{
	Clients map[*Client]bool
	Register chan *Client
	Unregister chan *Client
	Broadcast chan []byte
}

func (h *Hub) Run() {
	for {
		select {
		case Client := <-h.Register:
			h.Clients[Client] = true
		case Client := <-h.Unregister:
			h.Clients[Client] = false
	//	case message := h.Broadcast:
			// Loop through Clients and send message
		}
	}
}

func NewHub() *Hub{
	return &Hub{
		Clients: make(map[*Client]bool),
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast: make(chan []byte),
	}
}
