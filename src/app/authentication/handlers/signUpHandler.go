package handlers
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock_broker_application/constants"
	"stock_broker_application/models"
	"stock_broker_application/service"
)

// @Summary Sign up
// @Description Sign up with user details
// @Tags Authentication
// @Accept json
// @Produce json
// @Param signUpRequest body models.SignUpRequest true "Sign-up request payload"
// @Success 200 {object} string  "SuccessSignUp"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
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
		context.JSON(http.StatusOK, gin.H{"message": constants.SuccessSignUp})
	}
}
