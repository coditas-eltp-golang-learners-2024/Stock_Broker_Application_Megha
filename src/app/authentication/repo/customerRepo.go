package repo

import (
    "database/sql"
    "stock_broker_application/models"
    "stock_broker_application/utils/db"

)

// CheckCustomerExists checks if a customer already exists in the database
func CheckCustomerExists(phoneNumber, pancardNumber, password string) (bool, error) {
    var count int
    err := db.DB.QueryRow("SELECT COUNT(*) FROM customers WHERE phone_number = ? OR pancard_number = ? OR password = ?", phoneNumber, pancardNumber, password).Scan(&count)
    if err != nil && err != sql.ErrNoRows {
        return false, err
    }
    return count > 0, nil
}

// InsertCustomer inserts a new customer into the database
func InsertCustomer(customer *models.Customer) error {
    _, err := db.DB.Exec("INSERT INTO customer (name, email, phone_number, pancard_number, password) VALUES (?, ?, ?, ?, ?)", customer.Name, customer.Email, customer.PhoneNumber, customer.PancardNumber, customer.Password)
    return err
}
