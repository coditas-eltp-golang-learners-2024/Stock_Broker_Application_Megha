package repo

import (
    "database/sql"
    "stock_broker_application/models"
)

type CustomerRepository interface {
    CheckCustomerExists(phoneNumber, pancardNumber, password string) (bool, error)
    CheckCustomerExistsByEmailAndPassword(email, password string) (bool, error)
    InsertCustomer(customer *models.Customer) error
}

type CustomerDBRepository struct {
    DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
    return &CustomerDBRepository{DB: db}
}

func (r *CustomerDBRepository) CheckCustomerExists(phoneNumber, pancardNumber, password string) (bool, error) {
    var count int
    err := r.DB.QueryRow("SELECT COUNT(*) FROM customers WHERE phone_number = ? OR pancard_number = ? OR password = ?", phoneNumber, pancardNumber, password).Scan(&count)
    if err != nil && err != sql.ErrNoRows {
        return false, err
    }
    return count > 0, nil
}

func (r *CustomerDBRepository) CheckCustomerExistsByEmailAndPassword(email, password string) (bool, error) {
    var count int
    err := r.DB.QueryRow("SELECT COUNT(*) FROM customers WHERE email = ? AND password = ?", email, password).Scan(&count)
    if err != nil && err != sql.ErrNoRows {
        return false, err
    }
    return count > 0, nil
}

func (r *CustomerDBRepository) InsertCustomer(customer *models.Customer) error {
    _, err := r.DB.Exec("INSERT INTO customers (name, email, phone_number, pancard_number, password) VALUES (?, ?, ?, ?, ?)", customer.Name, customer.Email, customer.PhoneNumber, customer.PancardNumber, customer.Password)
    return err
}
