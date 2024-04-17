package models

// SignUpRequest represents the structure of user information.
type SignUpRequest struct {
	Name          string `gorm:"column:name;index" json:"name" validate:"required,min=3,max=50" example:"Megha Pawar"`
	Email         string `gorm:"column:email;uniqueIndex" json:"email" validate:"required,email,contains=@gmail.com"`
	PhoneNumber   string `gorm:"column:phone_number" json:"phone_number" validate:"required,min=10,max=10"`
	PancardNumber string `gorm:"column:pancard_number;uniqueIndex" json:"pancard_number" validate:"required,len=10"`
	Password      string `gorm:"column:password" json:"password" validate:"required,min=8" example:"password"`
}

// TableName sets the table name for SignUpRequest explicitly.
func (SignUpRequest) TableName() string {
	return "customers"
}
