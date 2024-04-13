package handlers

import (
    "net/http"
    "stock_broker_application/models"
    "stock_broker_application/service"
    "github.com/gin-gonic/gin"
    "stock_broker_application/constants"
)

// NewSignInHandler handles the sign-in request.
// @Summary Sign in
// @Description Sign in with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param email body string true "User email"
// @Param password body string true "User password"
// @Success 201 "Successful sign-in"
// @Failure 400 "Invalid request payload"
// @Failure 401 "Invalid credentials"
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
