package constants

import "errors"

var (
	ErrDatabaseInsert       = errors.New("failed to insert into database")
	ErrCustomerNotFound     = errors.New("customer not found")
	ErrInvalidCredentials   = errors.New("invalid Credentials")
	ErrDatabasePing         = errors.New("failed to ping database")
	ErrDatabaseConnection   = errors.New("failed to connect the database")
	ErrIncompleteFields     = errors.New("incomplete fields")
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidPhoneNumber   = errors.New("invalid phone number")
	ErrInvalidPancardNumber = errors.New("invalid pancard number")
	ErrMissingPassword      = errors.New("invalid password")
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrUnauthorized         = errors.New("unauthorized error")
	ErrCustAlreadyExists    = errors.New("customer already exists")
	ErrRecordNotFound       = errors.New("record not found")
	ErrEmailExists          = errors.New("email already exists")
	ErrPhoneNumberExists    = errors.New("phone number already exists")
	ErrInvalidPanCard       = errors.New("pancard number already exists")
	ErrInvalidPassword      = errors.New("invalid password")
)
