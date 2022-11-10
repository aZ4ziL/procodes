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
	// dsn := "fajhri:rafi213fajri@tcp(103.176.79.3:3306)/procodes?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:rafi213fajri@tcp(127.0.0.1:3306)/procodes?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
