package models

type SignInRequest struct {
	Email    string `gorm:"column:email" json:"email" validate:"required,email"`
	Password string `gorm:"column:password" json:"password" validate:"required,min=8"`
}

func (SignInRequest) TableName() string {
	return "customers"
}
