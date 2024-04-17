package repo

import (
	"gorm.io/gorm"
	"stock_broker_application/constants"
	"stock_broker_application/models"
)

// CustomerRepository defines the interface for interacting with customer data.
type CustomerRepository interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber string) bool
	IsPancardNumberExists(pancardNumber string) bool
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

func (repo *CustomerDBRepository) IsEmailExists(email string) bool {
	var count int64
	if err := repo.db.Model(&models.SignUpRequest{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *CustomerDBRepository) IsPhoneNumberExists(phoneNumber string) bool {
	var count int64
	if err := repo.db.Model(&models.SignUpRequest{}).Where("phone_number = ?", phoneNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *CustomerDBRepository) IsPancardNumberExists(pancardNumber string) bool {
	var count int64
	if err := repo.db.Model(&models.SignUpRequest{}).Where("pancard_number = ?", pancardNumber).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *CustomerDBRepository) InsertCustomer(customer *models.SignUpRequest) error {
	if err := repo.db.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CustomerDBRepository) GetUserByEmail(email string) (*models.SignInRequest, error) {
	var customer models.SignInRequest
	if err := repo.db.Where("email = ?", email).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, constants.ErrRecordNotFound
		}
		return nil, err
	}
	return &customer, nil
}
