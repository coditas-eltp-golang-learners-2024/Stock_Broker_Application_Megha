package db

import (
	"fmt"
	"log"
	"stock_broker_application/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// and pings the database to verify the connection.
func InitDB() {
	// Load database configuration
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading database configuration: %v", err)
	}

	// Construct Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.Username, config.Password, config.Host, config.Port, config.DBName)

	// Open a database connection
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Ping the database to verify the connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting DB instance: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	log.Println("Connected to database")
}

func GetDB() *gorm.DB {
	return DB
}
