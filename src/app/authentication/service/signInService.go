package service

import (
    "stock_broker_application/models"
    "errors"
)

type UserRepository interface {
    CheckCustomerExistsByEmailAndPassword(email, password string) (bool, error)
}

type SignInService struct {
    UserRepository UserRepository
}

func NewSignInService(userRepository UserRepository) *SignInService {
    return &SignInService{
        UserRepository: userRepository,
    }
}

func (s *SignInService) SignIn(signInRequest models.SignInRequest) error {
    exists, err := s.UserRepository.CheckCustomerExistsByEmailAndPassword(signInRequest.Email, signInRequest.Password)
    if err != nil {
        return err
    }
    if !exists {
        return errors.New("customer does not exist")
    }

    return nil
}
