package models
// SignUpRequest represents the structure of user information.
type SignUpRequest struct {
	Name          string `gorm:"column:name" json:"name" validate:"required,alpha"`
	Email         string `gorm:"column:email" json:"email" validate:"required,email"`
	PhoneNumber   string `gorm:"column:phone_number" json:"phone_number" validate:"required,gte=0000000000,lte=9999999999"`
	PancardNumber string `gorm:"column:pancard_number" json:"pancard_number" validate:"required,alphanum,len=10"`
	Password      string `gorm:"column:password" json:"password" validate:"required,alphanum,min=8"`
}
// TableName sets the table name for SignUpRequest explicitly.
func (SignUpRequest) TableName() string {
	return "customers"
}
