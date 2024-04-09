// models/signUpModel.go
package models

import (
"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
}

type Customer struct {
    Name          string `json:"name" validate:"required"`
    Email         string `json:"email" validate:"required,email"`
    PhoneNumber   string `json:"phone_number" validate:"required,len=10"`
    PancardNumber string `json:"pancard_number" validate:"required,len=10"`
    Password      string `json:"password" validate:"required,min=8"`
}

// Validate validates the Customer struct fields
func (c *Customer) Validate() error {
    return validate.Struct(c)
}
