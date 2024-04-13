// Package repo provides interfaces and implementations for interacting with customer data.
package repo

import (
	"stock_broker_application/models"
	"gorm.io/gorm"
)

// CustomerRepository defines the interface for interacting with customer data.
type CustomerRepository interface {
	
	CheckCustomerExists(phoneNumber, pancardNumber, password string) (bool, error)
	
	InsertCustomer(customer *models.SignUpRequest) error
	
	GetUserByEmail(email string) (*models.SignInRequest, error)
}

// CustomerDBRepository is an implementation of CustomerRepository using GORM for database interactions.
type CustomerDBRepository struct {
	db *gorm.DB
}

// NewCustomerRepository creates a new instance of CustomerDBRepository.
func NewCustomerRepository(db *gorm.DB) *CustomerDBRepository {
	return &CustomerDBRepository{db: db}
}

// It returns true if the customer exists, otherwise false.
func (repo *CustomerDBRepository) CheckCustomerExists(phoneNumber, pancardNumber, password string) (bool, error) {
	var count int64
	result := repo.db.Model(&models.SignUpRequest{}).Where("phone_number = ? OR pancard_number = ? OR password = ?", phoneNumber, pancardNumber, password).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (repo *CustomerDBRepository) InsertCustomer(customer *models.SignUpRequest) error {
	result := repo.db.Create(&customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CustomerDBRepository) GetUserByEmail(email string) (*models.SignInRequest, error) {
	var customer models.SignInRequest
	result := repo.db.Where("email = ?", email).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}
