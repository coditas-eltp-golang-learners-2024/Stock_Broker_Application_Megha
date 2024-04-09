package router

import (
    "github.com/gin-gonic/gin"
    "stock_broker_application/constants"
    "stock_broker_application/handlers"
    "stock_broker_application/service"
)

// SetupRouter initializes and sets up the router with dependencies injected.
func SetupRouter(signUpService *service.SignUpService) *gin.Engine {
    r := gin.Default()

    // Initialize SignUpHandler with SignUpService
    signUpHandler := handlers.NewSignUpHandler(signUpService)

    // Define route for signing up
    r.POST(constants.SignUpRoute, signUpHandler.SignUp)

    return r
}
