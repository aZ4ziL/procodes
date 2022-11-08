package controllers

import (
	"github.com/aZ4ziL/procodes/handlers"
	"github.com/gin-gonic/gin"
)

// ControllerAuthApiV1
func ControllerApiV1(router *gin.Engine) {
	apiAuth := handlers.NewApiV1Authentication()

	// Authentication
	v1 := router.Group("/api/v1")
	v1.POST("/register", apiAuth.Register())
	v1.POST("/auth", apiAuth.Auth())
	v1.GET("/logout", apiAuth.Logout())

	// BLog
	apiBlog := handlers.NewAPIV1Blog()
	v1.GET("/article", apiBlog.GetAll())
	v1.POST("/article", apiBlog.CreateArticle())
	v1.PUT("/article", apiBlog.UpdateArticle())
}
