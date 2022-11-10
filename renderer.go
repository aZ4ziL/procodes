package main

import (
	"text/template"

	"github.com/gin-contrib/multitemplate"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Template for Authentication Handler
	r.AddFromFiles("login", "views/users/login.tmpl")
	r.AddFromFiles("register", "views/users/register.tmpl")

	// Template for home
	r.AddFromFilesFuncs("index", template.FuncMap{},
		"views/home/index.tmpl")

	// Blog
	r.AddFromFilesFuncs("blog_index", template.FuncMap{
		"isPrime": isPrime,
	}, "views/blogs/base.tmpl", "views/blogs/index.tmpl")
	r.AddFromFilesFuncs("blog_detail", template.FuncMap{
		"markdown":             markdown,
		"getCategoryTitleByID": getBlogCategoryTitleByID,
	}, "views/blogs/base.tmpl", "views/blogs/detail.tmpl")

	// Chats
	r.AddFromFilesFuncs("chat_rooms", template.FuncMap{}, "views/chats/room.tmpl")

	return r
}
