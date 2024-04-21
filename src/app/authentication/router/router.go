// Package router provides functions to set up and configure the application's HTTP router.
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"stock_broker_application/constants"
	"stock_broker_application/docs"
	"stock_broker_application/handlers"
	"stock_broker_application/repo"
	"stock_broker_application/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Create a new Gin router with default middleware
	router := gin.Default()
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Create instances of the repository and service layers
	userRepository := repo.NewCustomerRepository(db)
	userService := service.NewSignUpService(userRepository)
	userAuthService := service.NewSignInService(userRepository)
	otpService := service.NewOTPService(userRepository)

	// Register HTTP routes for sign-up and sign-in endpoints
	router.POST(constants.SignUpRoute, handlers.NewSignUpHandler(userService))
	router.POST(constants.SignInRoute, handlers.NewSignInHandler(userAuthService))
	// router.POST(constants.OTPValidationRoute, handlers.NewValidateOTPHandler(otpService))
	router.POST("/validateOTP", handlers.NewValidateOTPHandler(otpService))

	return router
}
