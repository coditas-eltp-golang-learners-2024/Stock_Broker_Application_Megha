
package main

import (
	"log"
	"net/http"
	"stock_broker_application/router"
	"stock_broker_application/utils/db"
)

func main() {
    _, err := db.SetupDatabase()
    if err != nil {
        log.Fatal(err)
    }

    r:=router.SetupRouter()

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
