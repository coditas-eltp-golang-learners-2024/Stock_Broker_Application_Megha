package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_broker_application/constants"
	"stock_broker_application/models"
	"stock_broker_application/service"
)

// NewValidateOTPHandler handles the OTP validation request
// @Summary Validate OTP
// @Description Validates the OTP for a user
// @Tags OTP
// @Accept json
// @Produce json
// @Param otpRequest body models.OTPValidationRequest true "OTP Request"
// @Success 200 {string} string "OTP validated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "OTP is expired or invalid"
// @Router /validateOTP [post]
func NewValidateOTPHandler(otpService *service.OTPService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var otpValidationRequest models.OTPValidationRequest

		// Bind JSON request body to OTPValidationRequest struct
		if err := context.ShouldBindJSON(&otpValidationRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call OTPVerification method to validate OTP
		if err := otpService.OtpVerification(otpValidationRequest); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// OTP validation success
		context.JSON(http.StatusOK, gin.H{"message": constants.SuccessOTPValidation})
	}
}
