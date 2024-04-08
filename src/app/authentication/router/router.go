package router

import (
    "github.com/gin-gonic/gin"
    "stock_broker_application/constants"
    "stock_broker_application/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST(constants.SignUpRoute, handlers.SignUpHandler)

    return r
}
