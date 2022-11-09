package models

import (
	"time"

	"gorm.io/gorm"
)

type ChatGroup struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:100;unique" json:"title"`
	Members   []*User        `gorm:"many2many:user_chat_group_members" json:"members"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Create new chat group
func CreateNewChatGroup(group *ChatGroup) error {
	return db.Create(group).Error
}

// GetGroupByID get group by passing the ID
func GetGroupByID(id uint) (ChatGroup, error) {
	var group ChatGroup
	err := db.Model(&ChatGroup{}).Where("id = ?", id).First(&group).Preload("Members").Error
	return group, err
}

// GetAllChatGroups
func GetAllChatGroups() []ChatGroup {
	var groups []ChatGroup
	db.Model(&ChatGroup{}).Find(&groups).Preload("Members")
	return groups
}
