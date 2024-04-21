package service

import (
	"math/rand"
	"stock_broker_application/constants"
	"stock_broker_application/models"
	"stock_broker_application/repo"
	"time"
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
func (service *SignInService) SignIn(signInRequest models.SignInRequest) error {
	customers, err := service.UserRepository.GetUserByEmail(signInRequest.Email)

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

	otp, creationTime := generateOTP()
	if !service.UserRepository.AssignOtpToEmail(signInRequest.Email, otp, creationTime) {
		return constants.ErrOtpGeneration
	}
	return nil
}

func generateOTP() (int, time.Time) {
	creationTime := time.Now().Add(time.Minute)
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(10000)
	return otp, creationTime
}
