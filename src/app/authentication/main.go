package main

import (
    "log"
    "stock_broker_application/router"
    "stock_broker_application/utils/db"
)

func main() {
    // Initialize the database connection
    db.InitDB()
    // Defer closing the database connection
    // defer db.GetDB().Close()
    log.Println("Database connection is initialized, Application starting...")
    
    // Set up the router and start the application server
    r := router.SetupRouter(db.GetDB())
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
