
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

	dbConn, err := db.SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}

	customerRepo := repo.NewCustomerRepository(dbConn)

	signUpService := service.NewSignUpService(customerRepo)

	r := router.SetupRouter(signUpService)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
