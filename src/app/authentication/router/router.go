package router

import (
	"github.com/gin-gonic/gin"
	"stock_broker_application/constants"
	"stock_broker_application/handlers"
	"stock_broker_application/service"
)

func SetupRouter(signUpService *service.SignUpService, signInService *service.SignInService) *gin.Engine {
	r := gin.Default()

	// Initialize SignUpHandler with SignUpService
	signUpHandler := handlers.NewSignUpHandler(signUpService)
	r.POST(constants.SignUpRoute, signUpHandler.SignUp)

	// Initialize SignInHandler with SignInService
	signInHandler := handlers.NewSignInHandler(signInService)
	r.POST(constants.SignInRoute, signInHandler.SignIn)

	return r
}
