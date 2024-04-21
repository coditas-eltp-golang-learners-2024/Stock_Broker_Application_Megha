package repo

import (
	"github.com/go-sql-driver/mysql"
	"stock_broker_application/constants"
	"stock_broker_application/models"
	"time"
	// "database/sql"
	"gorm.io/gorm"
)

// CustomerRepository defines the interface for interacting with customer data.
type CustomerRepository interface {
	IsEmailExists(email string) bool
	IsPhoneNumberExists(phoneNumber string) bool
	IsPancardNumberExists(pancardNumber string) bool
	InsertCustomer(customer *models.SignUpRequest) error
	GetUserByEmail(email string) (*models.SignInRequest, error)

	CheckCredentialsExist(email, password string) bool
	AssignOtpToEmail(email string, otp int, creatime time.Time) bool
	CheckOtp(email string, otp int) bool
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

func (repo *CustomerDBRepository) CheckCredentialsExist(email, password string) bool {
	var count int64
	if err := repo.db.Model(&models.SignInRequest{}).
		Where("email = ? AND password = ?", email, password).
		Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *CustomerDBRepository) AssignOtpToEmail(email string, otp int, creationTime time.Time) bool {
	var count int64
	// Truncate milliseconds from otpExpiry
	creationTime = creationTime.Truncate(time.Second)
	if err := repo.db.Model(&models.OTPValidationRequest{}).Where("email = ?", email).Updates(models.OTPValidationRequest{OTP: otp, CreationTime: creationTime}).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (repo *CustomerDBRepository) CheckOtp(email string, otp int) bool {
	var count int64
	var otpCreationTime mysql.NullTime
	err := repo.db.Table("customers").Select("createdAt").Where("email = ?", email).Scan(&otpCreationTime).Error
	if err != nil {
		return false
	}
	if otpCreationTime.Valid {
		duration := time.Since(otpCreationTime.Time)
		if duration > 10*time.Minute {
			return false
		}
	}
	if err := repo.db.Model(&models.OTPValidationRequest{}).Where("email=? AND otp =?", email, otp).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}
