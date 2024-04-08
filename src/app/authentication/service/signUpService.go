package service

import (
	"database/sql"
	"errors"
	"stock_broker_application/models"
	"stock_broker_application/utils/db"
)

func SignUp(customer *models.Customer) error {
	if err := validateCustomer(customer); err != nil {
		return err
	}

	if err := checkDuplicateCustomer(customer); err != nil {
		return err
	}

	if err := insertCustomerIntoDB(customer); err != nil {
		return err
	}

	return nil
}

func validateCustomer(customer *models.Customer) error {
	if customer.Name == "" || customer.Email == "" || customer.PhoneNumber == "" || customer.PancardNumber == "" || customer.Password == "" {
		return errors.New("all fields are required")
	}

	return nil
}

func checkDuplicateCustomer(customer *models.Customer) error {
	var existingCustomer models.Customer
	err := db.DB.QueryRow("SELECT id FROM customers WHERE email = ? OR phone_number = ?", customer.Email, customer.PhoneNumber).Scan(&existingCustomer.ID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existingCustomer.ID != 0 {
		return errors.New("customer already exists with the same email or phone number")
	}

	return nil
}

func insertCustomerIntoDB(customer *models.Customer) error {
	_, err := db.DB.Exec("INSERT INTO customers (name, email, phone_number, pancard_number, password) VALUES (?, ?, ?, ?, ?)",
		customer.Name, customer.Email, customer.PhoneNumber, customer.PancardNumber, customer.Password)
	if err != nil {
		return err
	}

	return nil
}
