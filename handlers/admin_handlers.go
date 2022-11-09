package handlers

import (
	"net/http"

	"github.com/aZ4ziL/procodes/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminHandler interface {
	Index() gin.HandlerFunc
}

type adminHandler struct{}

func NewAdminHandler() AdminHandler {
	return &adminHandler{}
}

// Index for admin handler
func (a adminHandler) Index() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get session
		session := sessions.Default(ctx)

		user := session.Get("user")

		if user == nil {
			http.Error(ctx.Writer, "Please Login First. Permission Danied.", http.StatusForbidden)
			return
		}

		userToken := GetSessionValue(user)["token"].(string)

		userClaims, err := NewApiV1Authentication().CheckToken(userToken)
		if err != nil {
			http.Error(ctx.Writer, "Please Login First. Permission Danied.", http.StatusForbidden)
			return
		}

		// If user not superuser
		if !userClaims.IsSuperuser {
			http.Error(ctx.Writer, "Your account is not superuser, permission danied.", http.StatusForbidden)
			return
		}

		categories, articles := models.GetAllBlogCategories(), models.GetAllBlogArticles()
		users := models.GetAllUsers()

		ctx.HTML(http.StatusOK, "admin_index", gin.H{
			"user":       userClaims,
			"categories": categories,
			"articles":   articles,
			"users":      users,
		})
	}
}
