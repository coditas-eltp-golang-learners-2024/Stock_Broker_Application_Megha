package service

import (
	"stock_broker_application/constants"
	"stock_broker_application/models"
	"stock_broker_application/repo"
)

type OTPService struct {
	UserRepository repo.CustomerRepository
}

func NewOTPService(userRepository repo.CustomerRepository) *OTPService {
	return &OTPService{
		UserRepository: userRepository,
	}
}
func (otpService *OTPService) OtpVerification(otpData models.OTPValidationRequest) error {
	if !otpService.UserRepository.CheckOtp(otpData.Email, otpData.OTP) {
		return constants.ErrOtpVerification
	}
	return nil
}
