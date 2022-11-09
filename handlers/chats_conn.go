package handlers

import "github.com/gorilla/websocket"

type Message struct {
	RoomID    string `json:"room_id"`
	From      string `json:"from"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{}
}

type connection struct {
	ws   *websocket.Conn
	send chan *Message
}
