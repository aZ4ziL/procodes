package models

import (
	"time"

	"gorm.io/gorm"
)

type BlogCategory struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `gorm:"size:100;unique;index" json:"title" validate:"required"`
	Slug      string         `gorm:"size:100;unique;index" json:"slug" validate:"required"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Articles  []BlogArticle  `gorm:"foreignKey:CategoryID" json:"articles"`
}

// CreateNewBlogCategory
// create new category of blog
func CreateNewBlogCategory(category *BlogCategory) error {
	err := db.Create(category).Error
	return err
}

// GetBlogCategoryBySlug
// return a category by passing the SLUG
func GetBlogCategoryBySlug(slug string) (BlogCategory, error) {
	var category BlogCategory
	err := db.Model(&BlogCategory{}).Where("slug = ?", slug).First(&category).Error
	return category, err
}
