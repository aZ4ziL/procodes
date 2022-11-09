package controllers

import (
	"github.com/aZ4ziL/procodes/handlers"
	"github.com/gin-gonic/gin"
)

func IndexController(router *gin.Engine) {
	indexHandler := handlers.NewIndexHandler()

	router.GET("/", indexHandler.Home())
}

func BlogController(router *gin.Engine) {
	blog := router.Group("/blog")

	blog.GET("/", handlers.BlogIndex)
	blog.GET("/detail/:slug", handlers.BlogDetail)
}
