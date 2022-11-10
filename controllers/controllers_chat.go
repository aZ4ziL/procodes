package controllers

import (
	"github.com/aZ4ziL/procodes/handlers"
	"github.com/gin-gonic/gin"
)

func ChatController(router *gin.Engine) {
	chat := router.Group("/chat")
	chatHandler := handlers.NewChatHandler()

	chat.GET("room/:roomid", chatHandler.Room())
}
