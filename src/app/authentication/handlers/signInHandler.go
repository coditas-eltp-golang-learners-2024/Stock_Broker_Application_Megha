package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_broker_application/models"
	"stock_broker_application/service"
)

// NewSignInHandler handles the sign-in request
// @Summary Handle sign-in request
// @Description Handle sign-in request and authenticate the user
// @Accept json
// @Produce json
// @Param request body models.SignInRequest true "Sign-in request body"
// @Success 200 {string} string "User authenticated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /signIn [post]
func NewSignInHandler(userService *service.SignInService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var signInRequest models.SignInRequest

		// Bind JSON request body to SignInRequest struct
		if err := context.ShouldBindJSON(&signInRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call SignIn method to authenticate user
		if err := userService.SignIn(signInRequest); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
			return
		}

		// Authentication successful
		context.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
	}
}
