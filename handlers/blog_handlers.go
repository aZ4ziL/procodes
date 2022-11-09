package handlers

import (
	"fmt"
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
	categories := models.GetAllBlogCategories()

	if user != nil {
		userToken := GetSessionValue(user)["token"].(string)
		userClaims, err := NewApiV1Authentication().CheckToken(userToken)
		if err != nil {
			log.Println("user not logined")
		}
		ctx.HTML(http.StatusOK, "blog_index", gin.H{
			"user":       userClaims,
			"articles":   articles,
			"categories": categories,
		})
		return
	}

	ctx.HTML(http.StatusOK, "blog_index", gin.H{
		// "user":     userClaims,
		"articles":   articles,
		"categories": categories,
	})
}

// BlogDetail
func BlogDetail(ctx *gin.Context) {
	session := sessions.Default(ctx)

	slug := ctx.Param("slug")
	article, err := models.GetArticleBySlug(slug)
	categories := models.GetAllBlogCategories()

	if err != nil {
		http.Error(ctx.Writer, fmt.Sprintf("Article with slug `%s` not found error.", slug), http.StatusNotFound)
		return
	}

	user := session.Get("user")
	if user != nil {
		userToken := GetSessionValue(user)["token"].(string)
		userClaims, err := NewApiV1Authentication().CheckToken(userToken)
		if err != nil {
			log.Println("User not logined")
		}

		ctx.HTML(http.StatusOK, "blog_detail", gin.H{
			"user":       userClaims,
			"article":    article,
			"categories": categories,
		})
		return
	}

	ctx.HTML(http.StatusOK, "blog_detail", gin.H{
		"article":    article,
		"categories": categories,
	})
}
