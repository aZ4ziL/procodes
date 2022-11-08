package handlers

import "github.com/golang-jwt/jwt/v4"

// Declare secret key
var secretKey = []byte("mysecretkey")

type Credentials struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Claims struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsSuperuser bool   `json:"is_superuser"`
	jwt.RegisteredClaims
}
