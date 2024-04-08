// utils/db/sqlSetup.go

package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func SetupDatabase() (*sql.DB, error) {
	connectionString := "root:password@tcp(127.0.0.1:3306)/application"

	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")

	DB = database

	return DB, err
}
