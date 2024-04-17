package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_broker_application/constants"
	"stock_broker_application/models"
	"stock_broker_application/service"
)

// NewSignInHandler handles the sign-in request.
// @Summary Sign in
// @Description Sign in with user credentials
// @Tags Authentication
// @Accept json
// @Produce json
// @Param signInRequest body models.SignInRequest true "Sign-in request payload"
// @Success 201 {object} string "SuccessSignIn"
// @Failure 400 {object} string "Authentication failed"
// @Failure 401 {object} string "Invalid credentials"
// @Router /signIn [post]
func NewSignInHandler(signInService *service.SignInService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var signInRequest models.SignInRequest

		if err := context.BindJSON(&signInRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrAuthenticationFailed.Error()})
			return
		}
		// Call the SignIn method of SignInService to validate the sign-in request
		if err := signInService.SignIn(signInRequest); err != nil {
			// If there's an error, return the error response
			context.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrInvalidCredentials.Error()})
			return
		}
		// If there's no error, return the success message
		context.JSON(http.StatusCreated, gin.H{"message": constants.SuccessSignIn})
	}
}
