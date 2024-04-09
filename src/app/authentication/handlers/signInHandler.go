// handlers/signInHandler.go

package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "stock_broker_application/models"
    "stock_broker_application/service"
    "stock_broker_application/utils"
)

type SignInHandler struct {
    SignInService *service.SignInService
}

func NewSignInHandler(signInService *service.SignInService) *SignInHandler {
    return &SignInHandler{
        SignInService: signInService,
    }
}

func (h *SignInHandler) SignIn(c *gin.Context) {
    var signInRequest models.SignInRequest
    if err := c.BindJSON(&signInRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate the sign-in request
    if err := utils.ValidateSignInRequest(signInRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Convert SignInRequest to Customer
    customer := models.Customer{
        Email: signInRequest.Email,
        Password: signInRequest.Password,
    }

    // Call the SignIn method of the SignInService
    if err := h.SignInService.SignIn(customer); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User signed in successfully"})
}
