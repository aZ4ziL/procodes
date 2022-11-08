package models

import (
	"database/sql"
	"time"

	"github.com/aZ4ziL/procodes/auth"
)

type User struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	FirstName   string        `gorm:"size:100" json:"first_name" validate:"required"`
	LastName    string        `gorm:"size:100" json:"last_name" validate:"required"`
	Username    string        `gorm:"size:100;unique;index" json:"username" validate:"required"`
	Email       string        `gorm:"size:100;unique;index" json:"email" validate:"required, email"`
	Password    string        `gorm:"size:255" validate:"required" json:"password"`
	IsSuperuser bool          `gorm:"default:0" json:"is_superuser"`
	IsActive    bool          `gorm:"default:1" json:"is_active"`
	LastLogin   sql.NullTime  `json:"last_login"`
	DateJoined  time.Time     `gorm:"autoCreateTime" json:"date_joined"`
	Articles    []BlogArticle `gorm:"foreignKey:AuthorID" json:"articles"`
}

// CreateNewUser
// create new user by passing user struct
func CreateNewUser(user *User) error {
	user.Password = auth.EncryptionPassword(user.Password)
	err := db.Create(user).Error
	return err
}

// GetUserByUsername
// return a user by passing the USERNAME
func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.Model(&User{}).Where("username = ?", username).First(&user).Error
	return user, err
}

// GetUserByID
// return a user by passing the ID
func GetUserByID(id uint) (User, error) {
	var user User
	err := db.Model(&User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

// GetUserByEmail
// return a user by passing the EMAIL
func GetUserByEmail(email string) (User, error) {
	var user User
	err := db.Model(&User{}).Where("email = ?", email).First(&user).Error
	return user, err
}
