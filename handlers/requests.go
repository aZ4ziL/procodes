package handlers

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

// UserRequest for binding form post and or json post
type UserRequest struct {
	FirstName string `json:"first_name" validate:"required" form:"first_name"`
	LastName  string `json:"last_name" validate:"required" form:"last_name"`
	Username  string `json:"username" validate:"required" form:"username"`
	Email     string `json:"email" validate:"required,email" form:"email"`
	Password1 string `json:"password1" validate:"required" form:"password1"`
	Password2 string `json:"password2" validate:"required" form:"password2"`
}

func GetSessionValue(s interface{}) map[string]interface{} {
	values, _ := s.(map[string]interface{})
	result := make(map[string]interface{})

	for k, v := range values {
		result[k] = v
	}
	return result
}
