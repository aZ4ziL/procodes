package main

import (
	"bytes"
	"log"

	"github.com/aZ4ziL/procodes/models"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func isPrime(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

// markdown render the markdown to html
func markdown(str string) string {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(str), &buf); err != nil {
		log.Println(err.Error())
	}

	return buf.String()
}

// getBlogCategoryTitleByID
func getBlogCategoryTitleByID(id uint) string {
	db := models.GetDB()

	var category models.BlogCategory

	db.Model(&models.BlogCategory{}).Where("id = ?", id).First(&category)
	return category.Title
}
