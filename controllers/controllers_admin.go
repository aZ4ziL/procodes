package controllers

import (
	"github.com/aZ4ziL/procodes/handlers"
	"github.com/gin-gonic/gin"
)

func ControllerAdmin(router *gin.Engine) {
	adminHandler := handlers.NewAdminHandler()

	admin := router.Group("/admin")

	admin.GET("/", adminHandler.Index())
}
