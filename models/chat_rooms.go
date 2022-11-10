package models

import (
	"time"

	"gorm.io/gorm"
)

type ChatRoom struct {
	ID        uint           `gorm:"primaryKey"`
	GroupID   uint           `gorm:"index"`
	SenderID  uint           `gorm:"index"`
	Text      string         `gorm:"type:longtext"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateChatRoom
// create new chat room
func CreateChatRoom(room *ChatRoom) error {
	return db.Create(room).Error
}

// GetAllChatRooms
// Get all of chat rooms
func GetAllChatRooms() []ChatRoom {
	var rooms []ChatRoom
	db.Model(&ChatRoom{}).Find(&rooms)
	return rooms
}
