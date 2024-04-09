// models/signUpModel.go
package models
import(
    "stock_broker_application/utils"
)

type Customer struct {
    ID            int    `json:"id"`
    Name          string `json:"name" validate:"required"`
    Email         string `json:"email" validate:"required,email"`
    PhoneNumber   string `json:"phone_number" validate:"required,len=10"`
    PancardNumber string `json:"pancard_number" validate:"required,len=10"`
    Password      string `json:"password" validate:"required,min=8"`
}

func (c *Customer) Validate() error {
    return utils.ValidateStruct(c)
}
