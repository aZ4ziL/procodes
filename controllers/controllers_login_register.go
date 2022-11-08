package controllers

import (
	"github.com/aZ4ziL/procodes/handlers"
	"github.com/gin-gonic/gin"
)

func LoginRegisterLogout(router *gin.Engine) {
	router.GET("/login", handlers.Login)
	router.GET("/register", handlers.Register)
}
