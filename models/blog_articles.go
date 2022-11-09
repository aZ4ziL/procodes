package models

import (
	"time"

	"gorm.io/gorm"
)

type BlogArticle struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CategoryID uint           `json:"category_id" validate:"required"`
	AuthorID   uint           `json:"author_id" validate:"required"`
	Title      string         `gorm:"size:100;unique;index" json:"title" validate:"required"`
	Slug       string         `gorm:"size:100;unique;index" json:"slug" validate:"required"`
	Logo       string         `gorm:"size:255" json:"logo" validate:"required"`
	Desc       string         `gorm:"size:255" json:"desc" validate:"required"`
	Content    string         `gorm:"type:longtext" json:"content" validate:"required"`
	Status     string         `gorm:"size:9;default:DRAFTED"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CreateNewBlogArticle
// create new blog article
func CreateNewBlogArticle(article *BlogArticle) error {
	err := db.Create(article).Error
	return err
}

// GetArticleByID
func GetArticleByID(id uint) (BlogArticle, error) {
	var article BlogArticle
	err := db.Model(&BlogArticle{}).Where("id = ?", id).First(&article).Error
	return article, err
}

// GetArticleBySlug
// return article by passing the SLUG
func GetArticleBySlug(slug string) (BlogArticle, error) {
	var article BlogArticle
	err := db.Model(&BlogArticle{}).Where("slug = ?", slug).First(&article).Error
	return article, err
}

// GetAllBlogArticles
// return all articles with filter `status = PUBLISHED`
func GetAllBlogArticles() []BlogArticle {
	var articles []BlogArticle
	db.Model(&BlogArticle{}).Where("status = ?", "PUBLISHED").Find(&articles)
	return articles
}
