package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ChatHandler interface {
	Room() gin.HandlerFunc
}

type chatHandler struct{}

func NewChatHandler() ChatHandler {
	return &chatHandler{}
}

func (c chatHandler) Room() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		user := session.Get("user")

		if user == nil {
			http.Error(ctx.Writer, "Permission danied, please login first", http.StatusUnauthorized)
			return
		}

		userToken := GetSessionValue(user)["token"].(string)

		userClaims, err := NewApiV1Authentication().CheckToken(userToken)
		if err != nil {
			http.Error(ctx.Writer, "Please login first.", http.StatusUnauthorized)
			return
		}

		roomid := ctx.Param("roomid")

		ctx.HTML(http.StatusOK, "chat_rooms", gin.H{
			"user":   userClaims,
			"roomid": roomid,
		})
	}
}
