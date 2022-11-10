package handlers

import (
	"net/http"
	"strconv"

	"github.com/aZ4ziL/procodes/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ChatHandler interface {
	Index(ctx *gin.Context)
	Room(ctx *gin.Context)
}

type chatHandler struct{}

func NewChatHandler() ChatHandler {
	return &chatHandler{}
}

func (c chatHandler) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get("user")

	if user == nil {
		http.Error(ctx.Writer, "Please login first before accessing this page.", http.StatusUnauthorized)
		return
	}

	userToken := GetSessionValue(user)["token"].(string)
	userClaims, err := NewApiV1Authentication().CheckToken(userToken)
	if err != nil {
		http.Error(ctx.Writer, "Please login first before accessing this page.", http.StatusUnauthorized)
		return
	}

	// If user want to join to group chat
	if join, ok := ctx.GetQuery("join"); ok {
		// In the query join is category ID
		categoryID, _ := strconv.Atoi(join)

		if err != nil {
			http.Error(ctx.Writer, "Error while join to group. Please report this problem to web contact!", http.StatusBadRequest)
			return
		}
		err = models.AddChatGroupMember(userClaims.ID, uint(categoryID))
		if err != nil {
			http.Error(ctx.Writer, "Error while join to group. Please report this problem to web contact!", http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Successfully to joined group chat",
		})
		return
	}

	groups := models.GetAllChatGroups()
	ctx.HTML(http.StatusOK, "chat_index", gin.H{
		"groups": groups,
		"user":   userClaims,
	})
}

func (c chatHandler) Room(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get("user")
	if user == nil {
		http.Error(ctx.Writer, "Please login first before accessing this page.", http.StatusUnauthorized)
		return
	}

	userToken := GetSessionValue(user)["token"].(string)
	userClaims, err := NewApiV1Authentication().CheckToken(userToken)
	if err != nil {
		http.Error(ctx.Writer, "Please login first before accessing this page.", http.StatusUnauthorized)
		return
	}

	groupID := ctx.Param("groupID")
	groupIDInt, _ := strconv.Atoi(groupID)
	group, err := models.GetChatGroupByID(uint(groupIDInt))
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusNotFound)
		return
	}

	ctx.HTML(http.StatusOK, "chat_room", gin.H{
		"group": group,
		"user":  userClaims,
	})
}
