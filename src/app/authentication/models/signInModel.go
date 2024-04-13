
package models

// SignInCredentials represents the structure of user sign-in credentials.
type SignInRequest struct {
	Email    string `gorm:"column:email" json:"email" validate:"required,email"`
	Password string `gorm:"column:password" json:"password" validate:"required,min=8"`
}

// TableName sets the table name for SignInCredentials explicitly.
func (SignInRequest) TableName() string {
	return "customers"
}
