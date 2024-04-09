package service

import (
    "stock_broker_application/models"
    "errors"
)

type UserRepository interface {
    CheckCustomerExists(phoneNumber, pancardNumber, password string) (bool, error)
    InsertCustomer(customer *models.Customer) error
}

type SignUpService struct {
    UserRepository UserRepository
}

func NewSignUpService(userRepository UserRepository) *SignUpService {
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
