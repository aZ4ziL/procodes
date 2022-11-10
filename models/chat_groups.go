package models

import (
	"time"

	"gorm.io/gorm"
)

type ChatGroup struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string         `gorm:"size:100;unique"`
	AdminID   uint           `gorm:"index"`
	Members   []*User        `gorm:"many2many:user_chat_groups"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Rooms     []ChatRoom     `gorm:"foreignKey:GroupID"`
}

// CreateChatGroup
// Create new chat group
func CreateChatGroup(group *ChatGroup) error {
	return db.Create(group).Error
}

// GetChatGroupByID
// return chat group by passing the ID
func GetChatGroupByID(id uint) (ChatGroup, error) {
	var group ChatGroup
	err := db.Model(&ChatGroup{}).Where("id = ?", id).Preload("Members").Preload("Rooms").First(&group).Error
	return group, err
}

// GetAllChatGroups
// Get all of chat groups
func GetAllChatGroups() []ChatGroup {
	var groups []ChatGroup

	db.Model(&ChatGroup{}).Preload("Members").Find(&groups)
	return groups
}

func AddChatGroupMember(userID, groupID uint) error {
	user, err := GetUserByID(userID)
	if err != nil {
		return err
	}

	group, err := GetChatGroupByID(groupID)
	if err != nil {
		return err
	}

	db.Model(&group).Where("id = ?", groupID).Association("Members").Append([]*User{&user})
	return nil
}
