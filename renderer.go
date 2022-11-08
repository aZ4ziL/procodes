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

	return r
}
