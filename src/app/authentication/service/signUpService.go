package service
import (
    "errors"
    "stock_broker_application/models"
    "stock_broker_application/utils/db"
)

func SignUp(customer *models.Customer) error {
    // Validate the customer data
    if err := validateCustomer(customer); err != nil {
        return err
    }

    // Insert the customer into the database
    if err := insertCustomerIntoDB(customer); err != nil {
        return err
    }

    return nil
}

func validateCustomer(customer *models.Customer) error {
    // Implement validation logic for customer data here
    // For example:
    // Ensure that all fields are filled
    if customer.Name == "" || customer.Email == "" || customer.PhoneNumber == "" || customer.PancardNumber == "" || customer.Password == "" {
        return errors.New("all fields are required")
    }
    // You can add more validation checks for each field as per your requirements

    return nil
}

func insertCustomerIntoDB(customer *models.Customer) error {
    // Insert customer into the database using the db package
    _, err := db.DB.Exec("INSERT INTO customers (name, email, phone_number, pancard_number, password) VALUES (?, ?, ?, ?, ?)",
        customer.Name, customer.Email, customer.PhoneNumber, customer.PancardNumber, customer.Password)
    if err != nil {
        return err
    }

    return nil
}
