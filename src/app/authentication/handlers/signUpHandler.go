package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "stock_broker_application/models"
    "stock_broker_application/service"
    "stock_broker_application/constants"
)

// NewSignUpHandler handles the sign-up request.
// @Summary Sign up
// @Description Sign up with user details
// @Tags Authentication
// @Accept json
// @Produce json
// @Param signUpRequest body models.SignUpRequest true "Sign-up request payload"
// @Success 200 {object} models.SignUpRequest "Successful sign-up"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Internal server error"
// @Router /signUp [post]
func NewSignUpHandler(signUpService *service.SignUpService) gin.HandlerFunc {
    return func(context *gin.Context) {
        var signUpRequest models.SignUpRequest

        if err := context.BindJSON(&signUpRequest); err != nil {
            context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := signUpService.SignUp(signUpRequest); err != nil {
            context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        context.JSON(http.StatusCreated, gin.H{"message": constants.SuccessSignUp})
    }
}


