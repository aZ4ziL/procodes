package chatssockets

type Message struct {
	RoomID    string `json:"room_id"`
	From      string `json:"from"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

func NewMessage() *Message {
	return &Message{}
}
