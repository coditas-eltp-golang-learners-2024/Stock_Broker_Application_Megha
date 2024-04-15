package service

import (
	"stock_broker_application/constants"
	"stock_broker_application/models"
	"stock_broker_application/repo"
)

// SignInService handles sign-in logic
type SignInService struct {
	UserRepository repo.CustomerRepository
}

// NewSignInService creates a new instance of SignInService
func NewSignInService(userRepository repo.CustomerRepository) *SignInService {
	return &SignInService{
		UserRepository: userRepository,
	}
}

// SignIn authenticates the user
func (s *SignInService) SignIn(signInRequest models.SignInRequest) error {
	customers, err := s.UserRepository.GetUserByEmail(signInRequest.Email)

	if err != nil {
		// Handle the error if needed
		return err
	}

	if customers == nil {
		return constants.ErrCustomerNotFound
	}
	if customers.Password != signInRequest.Password {
		return constants.ErrInvalidCredentials
	}
	// Authentication successful
	return nil
}
