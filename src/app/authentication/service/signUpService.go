package service

import (
	"stock_broker_application/models"
	"stock_broker_application/repo"
)

// SignUpService handles user registration logic
type SignUpService struct {
	UserRepository repo.CustomerRepository
}

// NewSignUpService creates a new instance of SignUpService
func NewSignUpService(userRepository repo.CustomerRepository) *SignUpService {
	return &SignUpService{
		UserRepository: userRepository,
	}
}

// SignUp registers a new user with the provided details
func (s *SignUpService) SignUp(signUpRequest models.SignUpRequest) error {
	exists, err := s.UserRepository.CheckCustomerExists(signUpRequest.PhoneNumber, signUpRequest.PancardNumber, signUpRequest.Password)
	if err != nil {
		return err
	}
	if exists {
		return err
	}

	err = s.UserRepository.InsertCustomer(&signUpRequest)
	if err != nil {
		return err
	}

	return nil
}
