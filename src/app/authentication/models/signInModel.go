// models/signInModel.go

package models

import "github.com/go-playground/validator/v10"

type SignInRequest struct {
    Email       string `json:"email" validate:"required,email"`
    Password    string `json:"password" validate:"required,min=8"`
}

// Validate validates the SignInRequest struct fields
func (s *SignInRequest) Validate() error {
    validate := validator.New()
    return validate.Struct(s)
}
