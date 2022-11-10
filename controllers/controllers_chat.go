package controllers

import (
	"github.com/aZ4ziL/procodes/handlers"
	"github.com/aZ4ziL/procodes/handlers/chatssockets"
	"github.com/gin-gonic/gin"
)

func ChatController(router *gin.Engine) {
	chatHandler := handlers.NewChatHandler()

	hub := chatssockets.NewHub()
	go hub.Run()
	// socket
	router.GET("/ws/:roomId", func(ctx *gin.Context) {
		chatssockets.ServeWS(hub, ctx.Param("roomId"), ctx.Writer, ctx.Request)
	})

	chat := router.Group("/chat")

	chat.GET("/", chatHandler.Index)
	chat.GET("/:groupID", chatHandler.Room)
}
