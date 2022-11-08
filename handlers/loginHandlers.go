package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		ctx.HTML(http.StatusOK, "login", nil)
		return
	}
}

func Register(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		ctx.HTML(http.StatusOK, "register", nil)
		return
	}
}
