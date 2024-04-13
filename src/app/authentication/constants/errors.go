package constants

import "errors"

var (
	ErrDatabaseInsert     = errors.New("failed to insert into database")
	ErrCustomerNotFound   = errors.New("customer not found")
	ErrInvalidCredentials = errors.New("invalid Credentials")
	ErrDatabasePing         = errors.New("failed to ping database")
	ErrDatabaseConnection   = errors.New("failed to connect the database")
	ErrIncompleteFields     = errors.New("incomplete fields")
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPhoneNumber   = errors.New("invalid phone number")
	ErrInvalidPancardNumber = errors.New("invalid pancard number")
	ErrMissingPassword      = errors.New("invalid password")
	// SignInFailed=errors.New("sign in failed")
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrUnauthorized         = errors.New("unauthorized error")
)
