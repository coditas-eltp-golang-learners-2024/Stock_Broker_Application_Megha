package service

import (
	"stock_broker_application/constants"
	"stock_broker_application/models"
	"stock_broker_application/repo"
	"stock_broker_application/utils"
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
// SignUp validates and creates a new customer account.
func (service *SignUpService) SignUp(signUpRequest models.SignUpRequest) error {
	// Validate the SignUpRequest struct
	if err := utils.ValidateCustomer(signUpRequest); err != nil {
		return err
	}
	// Check if email already exists
	if service.UserRepository.IsEmailExists(signUpRequest.Email) {
		return constants.ErrEmailExists
	}

	// Check if phone number already exists
	if service.UserRepository.IsPhoneNumberExists(signUpRequest.PhoneNumber) {
		return constants.ErrInvalidPhoneNumber
	}

	// Check if pancard number already exists
	if service.UserRepository.IsPancardNumberExists(signUpRequest.PancardNumber) {
		return constants.ErrInvalidPanCard
	}

	if err := service.UserRepository.InsertCustomer(&signUpRequest); err != nil {
		return err
	}

	return nil
}
