package utils

import (
    "stock_broker_application/models"
    "github.com/go-playground/validator/v10"
    "stock_broker_application/constants"
)

func ValidateSignUpRequest(req models.SignUpRequest) error {
    // Create a new validator instance
    validate := validator.New()
    
    // Validate the name field
    if err := validate.Var(req.Name, "required,isAlpha"); err != nil {
        return constants.ErrInvalidName
    }

    // Validate the email field
    if err := validate.Var(req.Email, "required,email"); err != nil {
        return constants.ErrInvalidEmail
    }

    // Validate the phone number field
    if err := validate.Var(req.PhoneNumber, "required,len=10"); err != nil {
        return constants.ErrInvalidPhoneNumber
    }

    // Validate the PAN card number field
    if err := validate.Var(req.PancardNumber, "required,isAlphaNumeric"); err != nil {
        return constants.ErrInvalidPancardNumber
    }

    // Validate the password field
    if err := validate.Var(req.Password, "required"); err != nil {
        return constants.ErrMissingPassword
    }

    return nil
}
