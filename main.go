package main

import (
	"encoding/gob"
	"time"

	"github.com/aZ4ziL/procodes/controllers"
	"github.com/aZ4ziL/procodes/handlers"
	"github.com/aZ4ziL/procodes/models"
	"github.com/gin-contrib/sessions"
	gormsession "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func init() {
	gob.Register(time.Time{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[string]string{})
}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Static("/static", "./static")
	router.Static("/media", "./media")
	router.HTMLRender = createMyRender()

	// Set Gorm Session
	store := gormsession.NewStore(models.GetDB(), true, []byte("SecretKey"))
	router.Use(sessions.Sessions("procodesSessionID", store))

	router.SetTrustedProxies([]string{"127.0.0.1"})

	hub := handlers.NewHub()

	go hub.Run()

	router.GET("/ws/:roomid", func(ctx *gin.Context) {
		roomID := ctx.Param("roomid")
		handlers.ServeWS(hub, roomID, ctx.Writer, ctx.Request)
	})

	controllers.ControllerApiV1(router)
	controllers.IndexController(router)
	controllers.LoginRegisterLogout(router)
	controllers.BlogController(router)
	controllers.ChatController(router)

	router.Run(":8000")
}
