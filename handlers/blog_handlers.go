package handlers

import (
	"log"
	"net/http"

	"github.com/aZ4ziL/procodes/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// BlogIndex
func BlogIndex(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get("user")

	articles := models.GetAllBlogArticles()

	if user != nil {
		userToken := GetSessionValue(user)["token"].(string)
		userClaims, err := NewApiV1Authentication().CheckToken(userToken)
		if err != nil {
			log.Println("user not logined")
		}
		ctx.HTML(http.StatusOK, "blog_index", gin.H{
			"user":     userClaims,
			"articles": articles,
		})
		return
	}

	ctx.HTML(http.StatusOK, "blog_index", gin.H{
		// "user":     userClaims,
		"articles": articles,
	})
}
