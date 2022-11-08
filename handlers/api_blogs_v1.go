package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aZ4ziL/procodes/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type APIV1Blog interface {
	CreateArticle() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	UpdateArticle() gin.HandlerFunc
}

type apiV1Blog struct{}

// NewAPIV1Blog return instance of blog handlers
func NewAPIV1Blog() APIV1Blog {
	return &apiV1Blog{}
}

// CreateArticle handler for create new article
func (a apiV1Blog) CreateArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		user := session.Get("user")
		if user == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "unauthorized",
				"message": "Please login first before accessing this method.",
			})
			return
		}

		// Get Token user of user session
		userToken := GetSessionValue(user)["token"].(string)
		userClaims, err := NewApiV1Authentication().CheckToken(userToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "unauthorized",
				"message": "Please login first before accessing this method.",
			})
			return
		}

		// Check user is superuser or not
		// If user is not a superuser, return permission danied
		if !userClaims.IsSuperuser {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  "permission_danied",
				"message": "You is not a superuser, permission danied",
			})
			return
		}

		categoryID := ctx.PostForm("category_id")
		authorID := ctx.PostForm("author_id")
		title := ctx.PostForm("title")
		slug := ctx.PostForm("slug")
		desc := ctx.PostForm("desc")
		content := ctx.PostForm("content")
		status := ctx.PostForm("status")

		// Format String to int
		categoryIDInt, _ := strconv.Atoi(categoryID)
		authorIDInt, _ := strconv.Atoi(authorID)

		// Get File by form file name `logo`
		h, err := ctx.FormFile("logo")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Please upload file logo.",
			})
			return
		}

		// uuid get random uuid for filename
		filename := "media/articles/" + uuid.NewString() + h.Filename

		article := models.BlogArticle{
			CategoryID: uint(categoryIDInt),
			AuthorID:   uint(authorIDInt),
			Title:      title,
			Slug:       slug,
			Logo:       "/" + filename,
			Desc:       desc,
			Content:    content,
			Status:     status,
		}

		// validate
		validate = validator.New()
		err = validate.Struct(&article)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				log.Println(err.Error())
				return
			}

			errorMessages := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, with error type: %s", err.Field(), err.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": errorMessages,
			})
			return
		}

		// If no error save it
		err = models.CreateNewBlogArticle(&article)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		// Save file
		if err := ctx.SaveUploadedFile(h, filename); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Successfully to create new article with title = `" + article.Title,
		})
	}
}

// GetAll return all articles
func (a apiV1Blog) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// GetBySlug return article by given param SLUG
		if slug, ok := ctx.GetQuery("slug"); ok {
			article, err := models.GetArticleBySlug(slug)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status":  "error",
					"message": fmt.Sprintf("Article with slug `%s` is not found error.", slug),
				})
				return
			}

			ctx.JSON(http.StatusOK, article)
			return
		}

		articles := models.GetAllBlogArticles()
		ctx.JSON(http.StatusOK, articles)
	}
}

// UpdateArticle handler to update article using method PUT
func (a apiV1Blog) UpdateArticle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		// user session
		user := session.Get("user")
		if user == nil {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

		userToken := GetSessionValue(user)["token"].(string)

		userClaims, err := NewApiV1Authentication().CheckToken(userToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

		// is user is not superuser return error
		if !userClaims.IsSuperuser {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

		categoryID := ctx.PostForm("category_id")
		authorID := ctx.PostForm("author_id")
		title := ctx.PostForm("title")
		slug := ctx.PostForm("slug")
		desc := ctx.PostForm("desc")
		content := ctx.PostForm("content")
		status := ctx.PostForm("status")

		h, err := ctx.FormFile("logo")
		if err != nil {
			log.Println("logo is blank")
		}

		filename := "media/articles/" + uuid.NewString() + h.Filename

		// Parse string to int
		categoryIDInt, _ := strconv.Atoi(categoryID)
		authorIDInt, _ := strconv.Atoi(authorID)

		article, err := models.GetArticleBySlug(slug)
		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
			return
		}

		// Remove old file logo
		err = os.Remove("." + article.Logo)
		if err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		article.CategoryID = uint(categoryIDInt)
		article.AuthorID = uint(authorIDInt)
		article.Title = title
		article.Slug = slug
		article.Logo = "/" + filename
		article.Desc = desc
		article.Content = content
		article.Status = status

		if err := ctx.SaveUploadedFile(h, filename); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		models.GetDB().Save(&article)

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Successfully to update article.",
		})
	}
}
