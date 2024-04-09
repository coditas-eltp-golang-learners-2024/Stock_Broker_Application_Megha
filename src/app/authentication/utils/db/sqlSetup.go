package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDatabase() (*sql.DB, error) {
	connectionString := "root:password@tcp(127.0.0.1:3306)/application"

	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database!")

	return database, nil
}
