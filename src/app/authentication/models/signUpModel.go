// signUpModel.go

package models

import (
    _ "github.com/go-playground/validator/v10" // Explicitly import the package to inform the linter
)

type Customer struct {
    ID            int    `json:"id"`
    Name          string `json:"name" validate:"required"`
    Email         string `json:"email" validate:"required,email"`
    PhoneNumber   string `json:"phone_number" validate:"required,len=10"`
    PancardNumber string `json:"pancard_number" validate:"required,len=10"`
    Password      string `json:"password" validate:"required,min=8"`
}
