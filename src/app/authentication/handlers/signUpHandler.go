
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"stock_broker_application/models"
	"stock_broker_application/service"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func SignUpHandler(c *gin.Context) {
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if err := validate.Struct(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "validation_errors": err.Error()})
		return
	}

	
	err := service.SignUp(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer signed up successfully"})
}
