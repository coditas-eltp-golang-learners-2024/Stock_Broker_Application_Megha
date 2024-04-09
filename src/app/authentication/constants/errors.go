package constants

import "errors"

var (
	ErrValidationFailed = errors.New("validation failed")
	ErrCustomerExists   = errors.New("customer already exists")
	ErrInternalServer   = errors.New("internal server error")
)
