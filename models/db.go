package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = GetDB()

	db.AutoMigrate(&User{}, &BlogCategory{}, &BlogArticle{},
		&ChatGroup{}, &ChatRoom{})
}

func GetDB() *gorm.DB {
	dsn := "root:root@/procodes?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
