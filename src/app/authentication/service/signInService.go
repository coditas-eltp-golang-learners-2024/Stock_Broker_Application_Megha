
package service

import (
    "stock_broker_application/models"
    "errors"
)

type SignInService struct {
    UserRepository UserRepository
}

func NewSignInService(userRepository UserRepository) *SignInService {
    return &SignInService{
        UserRepository: userRepository,
    }
}

func (s *SignInService) SignIn(signInRequest models.Customer) error {
    exists, err := s.UserRepository.CheckCustomerExists(signInRequest.PhoneNumber, signInRequest.PancardNumber, signInRequest.Password)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("customer does not exist")
    }

    return nil
}
