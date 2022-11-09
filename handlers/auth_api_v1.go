package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aZ4ziL/procodes/auth"
	"github.com/aZ4ziL/procodes/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

// Login Authentication API version 1.0

type APIV1Authentication interface {
	Register() gin.HandlerFunc
	Auth() gin.HandlerFunc
	Logout() gin.HandlerFunc
	CheckToken(token string) (Claims, error)
}

type apiV1Authentication struct{}

func NewApiV1Authentication() APIV1Authentication {
	return &apiV1Authentication{}
}

// Register
// register user with method POST
func (a apiV1Authentication) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRequest UserRequest

		ctx.ShouldBindWith(&userRequest, binding.FormPost)

		validate = validator.New()

		// Validate the form
		err := validate.Struct(&userRequest)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				log.Println(err)
				return
			}
			errorMessages := []string{}
			for _, err := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s with wrong is: %s", err.Field(), err.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":   "error",
				"messages": errorMessages,
			})
			return
		}

		// Check password1 and password2 is same or not
		if userRequest.Password1 != userRequest.Password2 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Please reenter the password confirmation again.",
			})
			return
		}

		// Save it
		user := models.User{
			FirstName: userRequest.FirstName,
			LastName:  userRequest.LastName,
			Email:     userRequest.Email,
			Username:  userRequest.Username,
			Password:  userRequest.Password2,
		}

		err = models.CreateNewUser(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": fmt.Sprintf("Successfully to create new account with username `%s`", user.Username),
		})
	}
}

// Login
// Login user with Authentication JWT
func (a apiV1Authentication) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		session := sessions.Default(ctx)
		if user := session.Get("user"); user != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status":  "authorized",
				"message": "Already authorized.",
			})
			return
		}

		var creds Credentials

		// Get JSON FOrm with Form Data and decode into credentials
		err := ctx.ShouldBindWith(&creds, binding.FormPost)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Please enter the username and password."})
			return
		}

		user, err := models.GetUserByUsername(creds.Username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Username and password is incorrect."})
			return
		}

		// Check password with decryption method
		if !auth.DecryptionPassword(user.Password, creds.Password) {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Username and password is incorrect."})
			return
		}

		// Declare the expiration time of the token
		// here, we have kept it as 60 minutes
		expirationTime := time.Now().Add(30 * time.Hour)
		// Create JWT Claims, which includes the username and expired time
		claims := &Claims{
			ID:          user.ID,
			Username:    user.Username,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			IsSuperuser: user.IsSuperuser,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}
		// Declare the token with algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Error while decode the token.",
			})
			return
		}

		userSession := map[string]interface{}{
			"token": tokenString,
		}
		session.Set("user", userSession)
		if err := session.Save(); err != nil {
			http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "authorized",
			"message": "Successfully authorized.",
		})
	}
}

func (a apiV1Authentication) CheckToken(token string) (Claims, error) {
	var claims Claims

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tokenJWT, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return Claims{}, err
	}

	if !tokenJWT.Valid {
		return Claims{}, nil
	}

	return claims, nil
}

func (a apiV1Authentication) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		// Deleting all session
		session.Delete("user")
		session.Clear()
		session.Save()
		ctx.Redirect(http.StatusFound, "/login")
	}
}
