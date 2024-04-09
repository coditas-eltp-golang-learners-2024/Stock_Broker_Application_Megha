package utils

import (
    "github.com/go-playground/validator/v10"
    "stock_broker_application/models"
)

// Validator holds the validator instance
var Validator *validator.Validate

func init() {
    Validator = validator.New()
}

// ValidateSignUpRequest validates the sign-up request
func ValidateSignUpRequest(signUpRequest models.Customer) error {
    return Validator.Struct(signUpRequest)
}

// ValidateSignInRequest validates the sign-in request
func ValidateSignInRequest(signInRequest models.SignInRequest) error {
    return Validator.Struct(signInRequest)
}
