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
    if err := Validator.Struct(signUpRequest); err != nil {
        return err
    }
    return nil
}

// ValidateSignInRequest validates the sign-in request
func ValidateSignInRequest(signInRequest models.SignInRequest) error {
    if err := Validator.Struct(signInRequest); err != nil {
        return err
    }
    return nil
}
