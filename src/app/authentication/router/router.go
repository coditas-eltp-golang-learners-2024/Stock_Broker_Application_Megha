// Package router provides functions to set up and configure the application's HTTP router.
package router

import (
	"stock_broker_application/repo"
	"stock_broker_application/constants"
	"stock_broker_application/handlers"
	"stock_broker_application/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
    "github.com/swaggo/files" 
    "github.com/swaggo/gin-swagger" 
)

// SetupRouter initializes and configures the Gin HTTP router for the application.
// It takes a GORM database connection as input and returns the configured Gin engine.
func SetupRouter(db *gorm.DB) *gin.Engine {
	// Create a new Gin router with default middleware
	r := gin.Default()

	// Create instances of the repository and service layers
	userRepository := repo.NewCustomerRepository(db)
	userService := service.NewSignUpService(userRepository)
	userAuthService := service.NewSignInService(userRepository)

	// Register HTTP routes for sign-up and sign-in endpoints
	r.POST(constants.SignUpRoute, handlers.NewSignUpHandler(userService))
	r.POST(constants.SignInRoute, handlers.NewSignInHandler(userAuthService))

    r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
