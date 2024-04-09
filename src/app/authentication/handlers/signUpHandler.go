// handlers/signUpHandler.go
package handlers

import (
    "net/http"
	"github.com/gin-gonic/gin"
    "stock_broker_application/models"
    "stock_broker_application/service"
)

type SignUpHandler struct {
    SignUpService *service.SignUpService
}

func NewSignUpHandler(signUpService *service.SignUpService) *SignUpHandler {
    return &SignUpHandler{
        SignUpService: signUpService,
    }
}

func (h *SignUpHandler) SignUp(c *gin.Context) {
    var customer models.Customer
    if err := c.BindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := customer.Validate(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := h.SignUpService.SignUp(customer)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Customer signed up successfully"})
}
