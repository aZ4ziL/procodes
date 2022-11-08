package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type IndexHandler interface {
	Home() gin.HandlerFunc
}

type indexHandler struct{}

func NewIndexHandler() IndexHandler {
	return &indexHandler{}
}

func (i indexHandler) Home() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		user := session.Get("user")

		if user != nil {
			userToken := GetSessionValue(user)["token"].(string)

			claims, _ := NewApiV1Authentication().CheckToken(userToken)
			ctx.HTML(http.StatusOK, "index", gin.H{
				"user": claims,
			})
			return
		}

		ctx.HTML(http.StatusOK, "index", nil)
	}
}
