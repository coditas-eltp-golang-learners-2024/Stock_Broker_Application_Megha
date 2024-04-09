package service

import (
	"errors"
	"stock_broker_application/models"
	"stock_broker_application/repo"
)

type SignUpService struct {
	UserRepository repo.CustomerRepository
}

func NewSignUpService(userRepository repo.CustomerRepository) *SignUpService {
	return &SignUpService{
		UserRepository: userRepository,
	}
}

func (s *SignUpService) SignUp(signUpRequest models.Customer) error {
	exists, err := s.UserRepository.CheckCustomerExists(signUpRequest.PhoneNumber, signUpRequest.PancardNumber, signUpRequest.Password)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("customer already exists")
	}
	err = s.UserRepository.InsertCustomer(&signUpRequest)
	if err != nil {
		return err
	}
	return nil
}
