package handlers

type Hub struct {
	rooms      map[string]map[*connection]bool
	register   chan *subscription
	unregister chan *subscription
	broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		rooms:      make(map[string]map[*connection]bool),
		register:   make(chan *subscription),
		unregister: make(chan *subscription),
		broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.roomid]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.roomid] = connections
			}
			h.rooms[s.roomid][s.conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.roomid]
			if connections != nil {
				if _, ok := h.rooms[s.roomid]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.roomid)
					}
				}
			}
		case message := <-h.broadcast:
			connections := h.rooms[message.RoomID]
			for c := range connections {
				select {
				case c.send <- message:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, message.RoomID)
					}
				}
			}
		}
	}
}
