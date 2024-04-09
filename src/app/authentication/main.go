package main

import (
    "log"
    "net/http"
    "stock_broker_application/router"
    "stock_broker_application/utils/db"
    "stock_broker_application/repo"
    "stock_broker_application/service"
)

func main() {
    // Set up the database connection
    dbConn, err := db.SetupDatabase()
    if err != nil {
        log.Fatal(err)
    }

    // Create instances of the repositories
    customerRepo := repo.NewCustomerRepository(dbConn)

    // Create instances of the services with the repositories injected
    signUpService := service.NewSignUpService(customerRepo)
    signInService := service.NewSignInService(customerRepo)

    // Set up the router with the services injected
    r := router.SetupRouter(signUpService, signInService)
	
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
