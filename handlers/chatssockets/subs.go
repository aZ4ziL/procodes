package chatssockets

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type subscription struct {
	roomId string
	conn   *connection
	hub    *Hub
}

func (s *subscription) readPump() {
	c := s.conn
	defer func() {
		c.ws.Close()
		s.hub.unregister <- s
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(appData string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		message := NewMessage()
		err := c.ws.ReadJSON(message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				log.Println("error: ", err)
				return
			}
			break
		}
		message.RoomID = s.roomId
		message.Timestamp = time.Now().Format(time.Kitchen)
		s.hub.broadcast <- message
	}
}

func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.ws.WriteJSON(message)
		case <-ticker.C:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWS(hub *Hub, roomId string, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	c := &connection{
		ws:   ws,
		send: make(chan *Message, 256),
	}

	s := &subscription{
		conn:   c,
		hub:    hub,
		roomId: roomId,
	}

	s.hub.register <- s

	go s.writePump()
	go s.readPump()
}
